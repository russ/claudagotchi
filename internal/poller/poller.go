// Package poller runs an SSH probe against a remote host and returns the
// state of every active Claude Code session it finds under
// ~/.claude/projects.
package poller

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// HostSpec describes how to reach a single host. Name is required and used
// as the SSH target alias (which may resolve via ~/.ssh/config). The other
// fields, when non-empty, override the corresponding SSH options via -o
// flags on the command line.
type HostSpec struct {
	Name         string `toml:"name"`
	Hostname     string `toml:"hostname"`
	User         string `toml:"user"`
	Port         int    `toml:"port"`
	IdentityFile string `toml:"identity_file"`
}

// probeScript is sent on stdin to a remote `bash -s`. It enumerates the most
// recently modified top-level session jsonls and emits one SESSION line per
// session (TAB-separated, base64 last/prev to dodge quoting issues).
const probeScript = `set -u
PD="$HOME/.claude/projects"
echo "NOW=$(date +%s)"
if [ ! -d "$PD" ]; then
  echo "STATE=no_dir"
  exit 0
fi
echo "STATE=ok"
find "$PD" -mindepth 2 -maxdepth 2 -name "*.jsonl" -type f -printf "%T@\t%p\n" 2>/dev/null \
  | sort -rn | head -30 \
  | while IFS=$'\t' read -r MTIME FILE; do
      PROJECT=$(basename "$(dirname "$FILE")")
      UUID=$(basename "$FILE" .jsonl)
      LAST=$(tail -1 "$FILE" 2>/dev/null)
      PREV=$(tail -2 "$FILE" 2>/dev/null | head -1)
      LB64=$(printf '%s' "$LAST" | base64 | tr -d '\n')
      PB64=$(printf '%s' "$PREV" | base64 | tr -d '\n')
      printf 'SESSION\t%s\t%s\t%s\t%s\t%s\n' "$MTIME" "$PROJECT" "$UUID" "$LB64" "$PB64"
    done
`

// Phase classifies a single session's current activity.
type Phase int

const (
	PhaseUnknown Phase = iota
	PhaseUnreachable
	PhaseNoSessions
	PhaseSleeping
	PhaseIdle
	PhaseWaiting
	PhaseWorking
)

func (p Phase) String() string {
	switch p {
	case PhaseUnreachable:
		return "unreachable"
	case PhaseNoSessions:
		return "no sessions"
	case PhaseSleeping:
		return "sleeping"
	case PhaseIdle:
		return "idle"
	case PhaseWaiting:
		return "waiting"
	case PhaseWorking:
		return "working"
	}
	return "unknown"
}

// Session is one Claude Code session, identified by project + UUID.
type Session struct {
	SessionID         string
	Project           string
	ProjectPretty     string
	Phase             Phase
	SinceActive       time.Duration
	LastUserText      string
	LastAssistantText string
}

// Status is one host snapshot. If Err is non-empty the host is unreachable.
// Otherwise Sessions is the (possibly empty) list of recent sessions, sorted
// most-recent first.
type Status struct {
	Host     string
	Err      string
	Sessions []Session
	PolledAt time.Time
}

// FilterByAge returns the prefix of sessions whose SinceActive is no greater
// than window. Sessions are assumed sorted most-recent first. A non-positive
// window disables filtering.
func FilterByAge(ss []Session, window time.Duration) []Session {
	if window <= 0 {
		return ss
	}
	for i, s := range ss {
		if s.SinceActive > window {
			return ss[:i]
		}
	}
	return ss
}

func decodeProject(encoded string) string {
	s := strings.TrimPrefix(encoded, "-")
	s = strings.ReplaceAll(s, "-", "/")
	return "/" + s
}

func prettyProject(encoded string) string {
	p := decodeProject(encoded)
	parts := strings.Split(p, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-2] + "/" + parts[len(parts)-1]
	}
	return p
}

func extractText(line string) string {
	if line == "" {
		return ""
	}
	var msg struct {
		Type    string `json:"type"`
		Message struct {
			Role    string          `json:"role"`
			Content json.RawMessage `json:"content"`
		} `json:"message"`
	}
	if err := json.Unmarshal([]byte(line), &msg); err != nil {
		return truncate(line, 200)
	}
	if len(msg.Message.Content) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(msg.Message.Content, &s); err == nil {
		return s
	}
	var arr []struct {
		Type string `json:"type"`
		Text string `json:"text"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(msg.Message.Content, &arr); err == nil {
		for _, b := range arr {
			if b.Type == "text" && b.Text != "" {
				return b.Text
			}
			if b.Type == "tool_use" && b.Name != "" {
				return "🛠 " + b.Name
			}
		}
	}
	return ""
}

func messageRole(line string) string {
	if line == "" {
		return ""
	}
	var m struct {
		Type    string `json:"type"`
		Message struct {
			Role string `json:"role"`
		} `json:"message"`
	}
	_ = json.Unmarshal([]byte(line), &m)
	if m.Message.Role != "" {
		return m.Message.Role
	}
	return m.Type
}

func truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

func classify(age time.Duration, lastRole string) Phase {
	switch {
	case age < 5*time.Second:
		return PhaseWorking
	case age < 60*time.Second:
		if lastRole == "assistant" {
			return PhaseWaiting
		}
		return PhaseWorking
	case age < 5*time.Minute:
		return PhaseIdle
	default:
		return PhaseSleeping
	}
}

func parseSessionLine(line string, now int64) (Session, bool) {
	parts := strings.SplitN(line, "\t", 6)
	if len(parts) != 6 || parts[0] != "SESSION" {
		return Session{}, false
	}
	mtime, _ := strconv.ParseFloat(parts[1], 64)
	age := time.Duration(float64(now)-mtime) * time.Second
	if age < 0 {
		age = 0
	}

	var lastLine, prevLine string
	if b, err := base64.StdEncoding.DecodeString(parts[4]); err == nil {
		lastLine = string(b)
	}
	if b, err := base64.StdEncoding.DecodeString(parts[5]); err == nil {
		prevLine = string(b)
	}

	lastRole := messageRole(lastLine)
	prevRole := messageRole(prevLine)

	s := Session{
		SessionID:     parts[3],
		Project:       parts[2],
		ProjectPretty: prettyProject(parts[2]),
		SinceActive:   age,
		Phase:         classify(age, lastRole),
	}
	if lastRole == "user" {
		s.LastUserText = extractText(lastLine)
		if prevRole == "assistant" {
			s.LastAssistantText = extractText(prevLine)
		}
	} else {
		s.LastAssistantText = extractText(lastLine)
		if prevRole == "user" {
			s.LastUserText = extractText(prevLine)
		}
	}
	return s, true
}

// Poll runs the probe against host and returns a populated Status. It never
// returns an error; failures land in Status.Err.
func Poll(ctx context.Context, host HostSpec) Status {
	st := Status{Host: host.Name, PolledAt: time.Now()}
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	args := []string{
		"-T",
		"-o", "BatchMode=yes",
		"-o", "ConnectTimeout=5",
		"-o", "ControlMaster=auto",
		"-o", "ControlPath=~/.ssh/cm-%r@%h:%p",
		"-o", "ControlPersist=60",
	}
	if host.Hostname != "" {
		args = append(args, "-o", "HostName="+host.Hostname)
	}
	if host.User != "" {
		args = append(args, "-o", "User="+host.User)
	}
	if host.Port > 0 {
		args = append(args, "-o", fmt.Sprintf("Port=%d", host.Port))
	}
	if host.IdentityFile != "" {
		args = append(args, "-o", "IdentityFile="+host.IdentityFile)
	}
	args = append(args, host.Name, "bash -s")

	cmd := exec.CommandContext(cctx, "ssh", args...)
	cmd.Stdin = strings.NewReader(probeScript)
	var out, errb bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		msg := strings.TrimSpace(errb.String())
		if msg == "" {
			msg = err.Error()
		}
		st.Err = truncate(msg, 200)
		return st
	}

	fields := map[string]string{}
	var sessions []Session
	sc := bufio.NewScanner(&out)
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	// Two passes are not needed since NOW is emitted before any SESSION lines,
	// so we capture NOW first then parse SESSION lines as we read them.
	var nowEpoch int64
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "SESSION\t") {
			if s, ok := parseSessionLine(line, nowEpoch); ok {
				sessions = append(sessions, s)
			}
			continue
		}
		if i := strings.Index(line, "="); i > 0 {
			fields[line[:i]] = line[i+1:]
			if line[:i] == "NOW" {
				nowEpoch, _ = strconv.ParseInt(line[i+1:], 10, 64)
			}
		}
	}

	st.Sessions = sessions
	return st
}

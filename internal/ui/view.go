package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"

	"github.com/russ/claudagotchi/internal/poller"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("213")).
			Padding(0, 1)
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(1, 2)
	hostStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("75"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))
	quoteStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("248")).
			Italic(true)
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("105"))
	phaseStyle = map[poller.Phase]lipgloss.Style{
		poller.PhaseWorking:     lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true),
		poller.PhaseWaiting:     lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Bold(true),
		poller.PhaseIdle:        lipgloss.NewStyle().Foreground(lipgloss.Color("245")),
		poller.PhaseSleeping:    lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		poller.PhaseNoSessions:  lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		poller.PhaseUnreachable: lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true),
		poller.PhaseUnknown:     lipgloss.NewStyle().Foreground(lipgloss.Color("245")),
	}
)

const (
	// Inner panel chrome: rounded border (1 each side) + padding (2 each side).
	panelChrome   = 6
	minCardWidth  = 32
	maxCardCols   = 3
)

// View renders the full TUI frame.
func (m Model) View() string {
	if m.quitting {
		return ""
	}
	if m.width == 0 {
		return "loading…"
	}

	title := titleStyle.Render("✦  claudagotchi  ✦") +
		helpStyle.Render("  "+time.Now().Format("15:04:05"))

	n := len(m.hosts)
	if n == 0 {
		return title + "\nno hosts configured"
	}
	panelW := m.width/n - panelChrome
	if panelW < 32 {
		panelW = 32
	}

	panels := make([]string, 0, n)
	for _, h := range m.hosts {
		panels = append(panels, m.renderHost(h, panelW))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, panels...)

	help := helpStyle.Render(fmt.Sprintf("q quit · r refresh · poll every %s", m.pollInterval))

	return lipgloss.JoinVertical(lipgloss.Left, title, row, help)
}

// projectGroup collapses multiple sessions in the same project into one
// avatar card. Phase / SinceActive / messages come from the most recent
// session in the group (the input list is sorted most-recent first).
type projectGroup struct {
	Project           string
	ProjectPretty     string
	Count             int
	Phase             poller.Phase
	SinceActive       time.Duration
	LastUserText      string
	LastAssistantText string
}

func groupByProject(ss []poller.Session) []projectGroup {
	idx := map[string]int{}
	var groups []projectGroup
	for _, s := range ss {
		if i, ok := idx[s.Project]; ok {
			groups[i].Count++
			continue
		}
		idx[s.Project] = len(groups)
		groups = append(groups, projectGroup{
			Project:           s.Project,
			ProjectPretty:     s.ProjectPretty,
			Count:             1,
			Phase:             s.Phase,
			SinceActive:       s.SinceActive,
			LastUserText:      s.LastUserText,
			LastAssistantText: s.LastAssistantText,
		})
	}
	return groups
}

func sessionCount(gs []projectGroup) int {
	n := 0
	for _, g := range gs {
		n += g.Count
	}
	return n
}

func (m Model) renderHost(host string, panelW int) string {
	innerW := panelW
	if innerW < 24 {
		innerW = 24
	}

	s, polled := m.statuses[host]

	header := hostStyle.Render(host)
	var headerSuffix, content string

	inWindow := poller.FilterByAge(s.Sessions, m.activeWindow)
	groups := groupByProject(inWindow)
	hiddenGroups := 0
	if m.maxSessions > 0 && len(groups) > m.maxSessions {
		hiddenGroups = len(groups) - m.maxSessions
		groups = groups[:m.maxSessions]
	}

	switch {
	case !polled:
		headerSuffix = helpStyle.Render("…")
		content = renderHostFallback(poller.PhaseUnknown, m.frame, innerW, "waking up…")
	case s.Err != "":
		headerSuffix = phaseStyle[poller.PhaseUnreachable].Render("UNREACHABLE")
		content = renderHostFallback(poller.PhaseUnreachable, m.frame, innerW, s.Err)
	case len(s.Sessions) == 0:
		headerSuffix = helpStyle.Render("no sessions")
		content = renderHostFallback(poller.PhaseNoSessions, m.frame, innerW, "no Claude sessions found")
	case len(groups) == 0:
		headerSuffix = helpStyle.Render(fmt.Sprintf("idle >%s", durLabel(m.activeWindow)))
		content = renderHostFallback(poller.PhaseSleeping, m.frame, innerW,
			fmt.Sprintf("no activity in the last %s", durLabel(m.activeWindow)))
	default:
		n := sessionCount(groups)
		suffix := fmt.Sprintf("%d session%s", n, plural(n))
		if hiddenGroups > 0 {
			suffix += fmt.Sprintf(" · +%d project%s", hiddenGroups, plural(hiddenGroups))
		}
		headerSuffix = helpStyle.Render(suffix)
		content = m.renderGroupGrid(host, groups, hiddenGroups, innerW)
	}

	head := header
	if headerSuffix != "" {
		head += "  " + headerSuffix
	}

	body := head + "\n\n" + content
	return panelStyle.Width(panelW).Render(body)
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

// durLabel formats a duration like "1h" or "30m" without the "0s" tail
// time.Duration's String() can leave behind.
func durLabel(d time.Duration) string {
	if d <= 0 {
		return "∞"
	}
	if d%time.Hour == 0 {
		return fmt.Sprintf("%dh", int(d/time.Hour))
	}
	if d%time.Minute == 0 {
		return fmt.Sprintf("%dm", int(d/time.Minute))
	}
	return d.String()
}

func renderHostFallback(phase poller.Phase, frame, w int, msg string) string {
	sprite := RenderHostSprite(phase, frame)
	centered := centerArt(sprite, w)
	if msg == "" {
		return centered
	}
	wrapped := wrap(msg, w)
	centeredMsg := lipgloss.NewStyle().
		Width(w).
		Align(lipgloss.Center).
		Render(helpStyle.Render(wrapped))
	return centered + "\n\n" + centeredMsg
}

func (m Model) renderGroupGrid(host string, groups []projectGroup, hidden, w int) string {
	cols := w / minCardWidth
	if cols < 1 {
		cols = 1
	}
	if cols > maxCardCols {
		cols = maxCardCols
	}
	cardW := w / cols

	var rows []string
	for i := 0; i < len(groups); i += cols {
		end := i + cols
		if end > len(groups) {
			end = len(groups)
		}
		cards := make([]string, 0, end-i)
		for j := i; j < end; j++ {
			cards = append(cards, m.renderCard(host, groups[j], cardW))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cards...))
	}
	out := strings.Join(rows, "\n\n")
	if hidden > 0 {
		out += "\n\n" + helpStyle.Render(fmt.Sprintf("+%d more", hidden))
	}
	return out
}

func (m Model) renderCard(host string, g projectGroup, w int) string {
	creature := PickCreature(host, g.Project, m.creatures)
	var sprite string
	if creature != nil {
		sprite = creature.RenderSprite(g.Phase, m.frame)
	}
	if sprite == "" {
		sprite = RenderHostSprite(poller.PhaseUnknown, m.frame)
	}
	infoW := w - spriteWidth - 2
	if infoW < 12 {
		infoW = 12
	}

	phaseLine := phaseStyle[g.Phase].Render(strings.ToUpper(g.Phase.String()))

	project := g.ProjectPretty
	if g.Count > 1 {
		project += fmt.Sprintf(" ×%d", g.Count)
	}
	projectLine := truncate(project, infoW)

	ageLine := helpStyle.Render(humanDur(g.SinceActive) + " ago")

	youLine := ""
	if g.LastUserText != "" {
		youLine = labelStyle.Render("you: ") + quoteStyle.Render(truncate(oneline(g.LastUserText), infoW-5))
	}
	claudeLine := ""
	if g.LastAssistantText != "" {
		claudeLine = labelStyle.Render("claude: ") + quoteStyle.Render(truncate(oneline(g.LastAssistantText), infoW-8))
	}

	right := strings.Join([]string{phaseLine, projectLine, ageLine, youLine, claudeLine}, "\n")
	body := lipgloss.JoinHorizontal(lipgloss.Top, sprite, "  ", right)
	return lipgloss.NewStyle().Width(w).Render(body)
}

func centerArt(art string, width int) string {
	pad := (width - spriteWidth) / 2
	if pad <= 0 {
		return art
	}
	prefix := strings.Repeat(" ", pad)
	lines := strings.Split(art, "\n")
	for i := range lines {
		lines[i] = prefix + lines[i]
	}
	return strings.Join(lines, "\n")
}

func wrap(s string, w int) string {
	r := []rune(s)
	if len(r) <= w {
		return s
	}
	var b strings.Builder
	for i := 0; i < len(r); i += w {
		end := i + w
		if end > len(r) {
			end = len(r)
		}
		if i > 0 {
			b.WriteRune('\n')
		}
		b.WriteString(string(r[i:end]))
	}
	return b.String()
}

// truncate cuts s so that it occupies at most n visible columns, accounting
// for double-width and zero-width runes. The trailing ellipsis is included
// in the budget, so the returned string is never wider than n.
func truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return runewidth.Truncate(s, n, "…")
}

func oneline(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "\n", " ")
}

// HumanDur is exposed for the --once mode in main.
func HumanDur(d time.Duration) string { return humanDur(d) }

func humanDur(d time.Duration) string {
	if d < time.Second {
		return "now"
	}
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%dh", int(d.Hours()))
	}
	return fmt.Sprintf("%dd", int(d.Hours()/24))
}

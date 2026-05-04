// claudagotchi — a cute TUI that watches Claude Code agents over SSH.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/russ/claudagotchi/internal/config"
	"github.com/russ/claudagotchi/internal/poller"
	"github.com/russ/claudagotchi/internal/ui"
)

// version is overwritten via -ldflags="-X main.version=…" at build time.
var version = "dev"

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "claudagotchi:", err)
		os.Exit(1)
	}
}

func run() error {
	var (
		cfgPath     = flag.String("config", "", "path to config file (default: search XDG/home/cwd)")
		once        = flag.Bool("once", false, "poll each host once, print plain text, then exit")
		showVersion = flag.Bool("version", false, "print version and exit")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `claudagotchi — watch remote Claude Code agents

Usage:
  claudagotchi [flags] [host...]

Flags:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Hosts may be passed as positional arguments to override the config.

Config file (TOML):
  hosts = ["host-a", "host-b"]
  poll_interval = "3s"
  max_sessions  = 6
  active_window = "1h"

`)
	}
	flag.Parse()

	if *showVersion {
		fmt.Println("claudagotchi", version)
		return nil
	}

	cfg, _, err := config.Load(*cfgPath)
	if err != nil {
		return err
	}
	if hosts := flag.Args(); len(hosts) > 0 {
		cfg.Hosts = hosts
	}
	if len(cfg.Hosts) == 0 {
		return fmt.Errorf("no hosts configured (pass as args or set hosts in config)")
	}

	if *once {
		return runOnce(cfg.Hosts, cfg.MaxSessions, cfg.ActiveWindow.Duration)
	}

	p := tea.NewProgram(
		ui.New(cfg.Hosts, cfg.PollInterval.Duration, cfg.MaxSessions, cfg.ActiveWindow.Duration),
		tea.WithAltScreen(),
	)
	_, err = p.Run()
	return err
}

type onceGroup struct {
	Project           string
	Count             int
	Phase             poller.Phase
	SinceActive       time.Duration
	LastUserText      string
	LastAssistantText string
}

func groupSessions(ss []poller.Session) []onceGroup {
	idx := map[string]int{}
	var out []onceGroup
	for _, s := range ss {
		if i, ok := idx[s.ProjectPretty]; ok {
			out[i].Count++
			continue
		}
		idx[s.ProjectPretty] = len(out)
		out = append(out, onceGroup{
			Project:           s.ProjectPretty,
			Count:             1,
			Phase:             s.Phase,
			SinceActive:       s.SinceActive,
			LastUserText:      s.LastUserText,
			LastAssistantText: s.LastAssistantText,
		})
	}
	return out
}

func runOnce(hosts []string, maxSessions int, activeWindow time.Duration) error {
	for _, h := range hosts {
		s := poller.Poll(context.Background(), h)
		fmt.Printf("=== %s ===\n", s.Host)
		total := len(s.Sessions)
		inWindow := poller.FilterByAge(s.Sessions, activeWindow)
		groups := groupSessions(inWindow)
		hiddenGroups := 0
		if maxSessions > 0 && len(groups) > maxSessions {
			hiddenGroups = len(groups) - maxSessions
			groups = groups[:maxSessions]
		}

		switch {
		case s.Err != "":
			fmt.Printf("unreachable: %s\n", s.Err)
		case total == 0:
			fmt.Println("no sessions")
		case len(groups) == 0:
			fmt.Printf("no activity in the last %s (%d older session%s)\n",
				activeWindow, total, pluralS(total))
		default:
			n := 0
			for _, g := range groups {
				n += g.Count
			}
			fmt.Printf("%d session%s in %d project%s:\n",
				n, pluralS(n), len(groups), pluralS(len(groups)))
			for _, g := range groups {
				project := g.Project
				if g.Count > 1 {
					project += fmt.Sprintf(" ×%d", g.Count)
				}
				fmt.Printf("  %-9s  %-30s  %s ago",
					g.Phase, project, ui.HumanDur(g.SinceActive))
				if g.LastUserText != "" {
					fmt.Printf("  · you: %s",
						truncate(strings.ReplaceAll(g.LastUserText, "\n", " "), 60))
				} else if g.LastAssistantText != "" {
					fmt.Printf("  · claude: %s",
						truncate(strings.ReplaceAll(g.LastAssistantText, "\n", " "), 60))
				}
				fmt.Println()
			}
			if hiddenGroups > 0 {
				fmt.Printf("  +%d more project%s\n", hiddenGroups, pluralS(hiddenGroups))
			}
		}
		fmt.Println()
	}
	return nil
}

func pluralS(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

func truncate(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

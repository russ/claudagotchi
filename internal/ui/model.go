// Package ui implements the Bubble Tea program: model, view, and sprites.
package ui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/russ/claudagotchi/internal/poller"
)

const animationInterval = 300 * time.Millisecond

type tickMsg time.Time
type animMsg struct{}
type pollResultMsg poller.Status

// Model is the Bubble Tea model. Construct with New.
type Model struct {
	hosts        []poller.HostSpec
	pollInterval time.Duration
	maxSessions  int
	activeWindow time.Duration
	creatures    []string
	statuses     map[string]poller.Status
	width        int
	height       int
	quitting     bool
	frame        int
}

// New returns a Model configured with the given hosts, poll interval, max
// sessions to display per host, the activity window beyond which sessions
// are hidden, and the creature roster (empty = all available).
func New(hosts []poller.HostSpec, pollInterval time.Duration, maxSessions int, activeWindow time.Duration, creatures []string) Model {
	if pollInterval <= 0 {
		pollInterval = 3 * time.Second
	}
	if maxSessions <= 0 {
		maxSessions = 6
	}
	return Model{
		hosts:        hosts,
		pollInterval: pollInterval,
		maxSessions:  maxSessions,
		activeWindow: activeWindow,
		creatures:    creatures,
		statuses:     map[string]poller.Status{},
	}
}

func (m Model) tickCmd() tea.Cmd {
	return tea.Tick(m.pollInterval, func(t time.Time) tea.Msg { return tickMsg(t) })
}

func animCmd() tea.Cmd {
	return tea.Tick(animationInterval, func(_ time.Time) tea.Msg { return animMsg{} })
}

func pollCmd(host poller.HostSpec) tea.Cmd {
	return func() tea.Msg {
		return pollResultMsg(poller.Poll(context.Background(), host))
	}
}

// Init starts the poll and animation loops.
func (m Model) Init() tea.Cmd {
	cmds := []tea.Cmd{m.tickCmd(), animCmd()}
	for _, h := range m.hosts {
		cmds = append(cmds, pollCmd(h))
	}
	return tea.Batch(cmds...)
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "r":
			cmds := make([]tea.Cmd, 0, len(m.hosts))
			for _, h := range m.hosts {
				cmds = append(cmds, pollCmd(h))
			}
			return m, tea.Batch(cmds...)
		}
	case tickMsg:
		cmds := []tea.Cmd{m.tickCmd()}
		for _, h := range m.hosts {
			cmds = append(cmds, pollCmd(h))
		}
		return m, tea.Batch(cmds...)
	case animMsg:
		m.frame++
		return m, animCmd()
	case pollResultMsg:
		m.statuses[msg.Host] = poller.Status(msg)
	}
	return m, nil
}

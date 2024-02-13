package main

// The combination of items from lipglossDynamic.go and lipglossPureProgress.go

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	timeout  = time.Second * 5
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type model struct {
	percent  float64
	progress progress.Model
	barType  string
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		switch m.barType {
		case "dynamic":
			if m.progress.Percent() == 1.0 {
				return m, tea.Quit
			}

			// Note that you can also use progress.Model.SetPercent to set the
			// percentage value explicitly, too.
			cmd := m.progress.IncrPercent(0.25)
			return m, tea.Batch(tickCmd(), cmd)

		case "pure":
			m.percent += 0.25
			if m.percent > 1.0 {
				m.percent = 1.0
				return m, tea.Quit
			}
			return m, tickCmd()

		default:
			return m, tickCmd()
		}

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m model) View() string {
	pad := strings.Repeat(" ", padding)
	switch m.barType {
	case "dynamic":
		return "\n" +
			pad + m.progress.View() + "\n\n" +
			pad + "And we are moving on...\n\n"
		// pad + helpStyle("Press any key to quit")
	case "pure":
		return "\n" +
			pad + m.progress.ViewAs(m.percent) + "\n\n" +
			pad + "And we are moving on...\n\n"
		// pad + helpStyle("Press any key to quit")
	default:
		return "\n" +
			pad + helpStyle("Press any key to quit")
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type timerModel struct {
	timer    timer.Model
	keymap   keymap
	help     help.Model
	quitting bool
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func (tm timerModel) Init() tea.Cmd {
	return tm.timer.Init()
}

func (tm timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		tm.timer, cmd = tm.timer.Update(msg)
		return tm, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		tm.timer, cmd = tm.timer.Update(msg)
		// Toggle whether 's' does s start or stop
		tm.keymap.stop.SetEnabled(tm.timer.Running())
		tm.keymap.start.SetEnabled(!tm.timer.Running())
		return tm, cmd

	case timer.TimeoutMsg:
		tm.quitting = true
		return tm, tea.Quit

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, tm.keymap.quit):
			tm.quitting = true
			return tm, tea.Quit
		case key.Matches(msg, tm.keymap.reset):
			tm.timer.Timeout = timeout
		case key.Matches(msg, tm.keymap.start, tm.keymap.stop):
			return tm, tm.timer.Toggle()
		}
	}

	return tm, nil
}

func (tm timerModel) helpView() string {
	return "\n" + tm.help.ShortHelpView([]key.Binding{
		tm.keymap.start,
		tm.keymap.stop,
		tm.keymap.reset,
		tm.keymap.quit,
	})
}

func (tm timerModel) View() string {
	// For a more detailed timer view you could read tm.timer.Timeout to get
	// the remaining time as a time.Duration and skip calling tm.timer.View()
	// entirely.
	s := tm.timer.View()

	if tm.timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	if !tm.quitting {
		s = "Exiting in " + s
		s += tm.helpView()
	}
	return s
}

func RunBubbleTimer() {
	tm := timerModel{
		timer: timer.NewWithInterval(timeout, time.Millisecond),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		},
		help: help.New(),
	}
	tm.keymap.start.SetEnabled(false)

	if _, err := tea.NewProgram(tm).Run(); err != nil {
		fmt.Println("Uh oh, we encountered an error:", err)
		return
	}
	
}

package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type tutorialModel struct {
	choices []string // todo items
	cursor int // todo item the cursor is on
	selected map[int]struct{}	// selected todo items
}

// Define todo list initial state
func initializeModel() tea.Model {
	return tutorialModel{
		// The todo list
		choices: []string{"Write Docs", "Close ticket", "Do golang practice"},

		// Map indicating which choices are selected.
		selected: make(map[int]struct{}),
}
}

// Create the `Cmd` for doing I/O. Not needed in this
// instance
func (tm tutorialModel) Init() tea.Cmd {
	// Returning nil means "no I/O now"
	return nil
}

// Update() is called when "things happen" (aka "a message is received").
// Update() looks at the message, updates the model as needed, and returns
// the updated model
func (tm tutorialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Find out what kind of message it is
	switch msg := msg.(type) {

		// Was a key pressed?
	case tea.KeyMsg:
		// if so, what key was it?
		switch msg.String() {
		case "ctrl+c", "q":
			return tm, tea.Quit

		case "up", "k":
			if tm.cursor > 0 {
				tm.cursor--
			}

		case "down", "j":
			if tm.cursor < len(tm.choices)-1 {
				tm.cursor++
			}

		case "enter", " ":
			_, ok := tm.selected[tm.cursor]
			if ok {
				delete(tm.selected, tm.cursor)
			} else {
				tm.selected[tm.cursor] = struct{}{}
			}
		}
	}

	// Now that you've updated the model, return it. We don't return a command
	// in this tutorial.
	return tm, nil
}

// Thing to render the model to the screen. This just returns a string and
// Bubble Tea does the magic behind the scenes to do appropriate redraws
func (tm tutorialModel) View() string {
	// Header
	s := "What should we do today, Brain?\n\n"

	for i, choice := range tm.choices {
		cursor := " "	// No cursor
		if tm.cursor == i {
			cursor = ">"	// magical cursor
		}

		// Is the choice selected?
		checked := " "	// not selected
		if _, ok := tm.selected[i]; ok {
			checked = "x" // selected
		}

		// "Render" the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// footer
	s += "\nPress q to quit.\n"

	// And return the string for the magic
	return s
}

// And what runs it all
func RunBubbleTutorial() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("We got an error. On noes: %v", err)
		return
	}
}
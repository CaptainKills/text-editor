package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Cursor struct {
	row    int
	column int
}

type Model struct {
	fileName string
	buffer   []string
	cursor   Cursor
	command  string
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		case "backspace":
			m.command = m.command[:max(0, len(m.command)-1)]

		case "enter", "esc":
			m.command = ""

		case "h", "left":
			m.cursor.column = max(m.cursor.column-1, 0)

		case "j", "down":
			m.cursor.row = min(m.cursor.row+1, len(m.buffer)-1)
			// TODO: Add column update

		case "k", "up":
			m.cursor.row = max(m.cursor.row-1, 0)
			// TODO: Add column update

		case "l", "right":
			m.cursor.column = min(m.cursor.column+1, len(m.buffer[m.cursor.row]))

		default:
			m.command += msg.String()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	editor := lipgloss.JoinHorizontal(
		lipgloss.Top,
		RenderNumberColumn(m),
		RenderCode(m),
	)

	editor = lipgloss.JoinVertical(
		lipgloss.Top,
		editor,
		RenderStatusLine(m),
		RenderCommandLine(m),
	)

	// Send the UI for rendering
	return editor
}

package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

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
			m.cursor.MoveLeft(m.buffer)

		case "j", "down":
			m.cursor.MoveDown(m.buffer)

		case "k", "up":
			m.cursor.MoveUp(m.buffer)

		case "l", "right":
			m.cursor.MoveRight(m.buffer)

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

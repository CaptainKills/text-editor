package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Normal  int = 0
	Insert  int = 1
	Command int = 2
	Visual  int = 3
)

type Model struct {
	fileName string

	mode    int
	buffer  []string
	cursor  Cursor
	command string

	log *os.File
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

		case "enter":
			m.command = ""

		// Movement Keys
		case "h", "left":
			m.cursor.MoveLeft(m.buffer)

		case "j", "down":
			m.cursor.MoveDown(m.buffer)

		case "k", "up":
			m.cursor.MoveUp(m.buffer)

		case "l", "right":
			m.cursor.MoveRight(m.buffer)

		// Mode Keys
		case "esc":
			m.mode = Normal
			m.command = ""

		case "i":
			m.mode = Insert

		case ":":
			m.mode = Command
			m.command = ":"

		case "v":
			m.mode = Visual

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
		RenderStatusColumn(m),
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

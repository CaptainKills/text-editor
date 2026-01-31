package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	NormalMode  int = 0
	InsertMode  int = 1
	CommandMode int = 2
	VisualMode  int = 3
	SearchMode  int = 4
)

var ModeString = map[int]string{
	NormalMode:  "Normal",
	InsertMode:  "Insert",
	CommandMode: "Command",
	VisualMode:  "Visual",
	SearchMode:  "Search",
}

type Model struct {
	fileName string

	mode    int
	buffer  []string
	cursor  Cursor
	command string

	log *os.File

	width  int
	height int
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Mode Agnostic Update Handling
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Fprintf(m.log, "(WINDOW) width: %d, height: %d\n", msg.Width, msg.Height)
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		// Program Exit
		case "ctrl+c", "q":
			return m, tea.Quit

		// Return to Normal Mode
		case "esc":
			m.mode = NormalMode
			m.command = ""
		}
	}

	// Mode Specific Update Handling
	switch m.mode {
	case NormalMode:
		cmd = m.normalModeUpdate(msg)
	case InsertMode:
		cmd = m.insertModeUpdate(msg)
	case CommandMode:
		cmd = m.commandModeUpdate(msg)
	case VisualMode:
		cmd = m.visualModeUpdate(msg)
	case SearchMode:
		cmd = m.searchModeUpdate(msg)
	default:
		cmd = m.normalModeUpdate(msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m Model) View() tea.View {
	ui := lipgloss.JoinHorizontal(
		lipgloss.Top,
		RenderStatusColumn(m),
		RenderCode(m),
	)

	ui = lipgloss.JoinVertical(
		lipgloss.Top,
		ui,
		RenderStatusLine(m),
		RenderCommandLine(m),
	)

	// Send the UI for rendering
	view := tea.NewView(ui)
	view.AltScreen = true

	return view
}

func (m *Model) normalModeUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
		case "i":
			m.mode = InsertMode

		case ":":
			m.mode = CommandMode
			m.command = ":"

		case "/":
			m.mode = SearchMode
			m.command = "/"

		case "v":
			m.mode = VisualMode
		}
	}

	return nil
}

func (m *Model) insertModeUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return nil
}

func (m *Model) commandModeUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			if len(m.command) <= 1 {
				m.mode = NormalMode
				m.command = ""
				break
			}
			m.command = m.command[:max(0, len(m.command)-1)]

		case "enter":
			m.mode = NormalMode
			m.command = ""

		default:
			m.command += msg.String()
		}
	}

	return nil
}

func (m *Model) visualModeUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}

	return nil
}

func (m *Model) searchModeUpdate(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "backspace":
			if len(m.command) <= 1 {
				m.mode = NormalMode
				m.command = ""
				break
			}
			m.command = m.command[:max(0, len(m.command)-1)]

		case "enter":
			m.mode = NormalMode
			m.command = ""

		default:
			m.command += msg.String()
		}
	}

	return nil
}

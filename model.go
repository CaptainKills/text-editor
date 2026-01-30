package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Normal  int = 0
	Insert  int = 1
	Command int = 2
	Visual  int = 3
	Search  int = 4
)

var ModeString = map[int]string{
	Normal:  "Normal",
	Insert:  "Insert",
	Command: "Command",
	Visual:  "Visual",
	Search:  "Search",
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
			m.mode = Normal
			m.command = ""
		}
	}

	// Mode Specific Update Handling
	switch m.mode {
	case Normal:
		cmd = m.normalModeUpdate(msg)
	case Insert:
		cmd = m.insertModeUpdate(msg)
	case Command:
		cmd = m.commandModeUpdate(msg)
	case Visual:
		cmd = m.visualModeUpdate(msg)
	case Search:
		cmd = m.searchModeUpdate(msg)
	default:
		cmd = m.normalModeUpdate(msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m Model) View() string {
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
	return ui
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
			m.mode = Insert

		case ":":
			m.mode = Command
			m.command = ":"

		case "/":
			m.mode = Search
			m.command = "/"

		case "v":
			m.mode = Visual
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
				m.mode = Normal
				m.command = ""
				break
			}
			m.command = m.command[:max(0, len(m.command)-1)]

		case "enter":
			m.mode = Normal
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
				m.mode = Normal
				m.command = ""
				break
			}
			m.command = m.command[:max(0, len(m.command)-1)]

		case "enter":
			m.mode = Normal
			m.command = ""

		default:
			m.command += msg.String()
		}
	}

	return nil
}

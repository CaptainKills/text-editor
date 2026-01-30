package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderStatusLine(m Model) string {
	mode := renderMode(m)
	filepath := renderFilePath(m)
	cursor := renderCursor(m)

	ui := lipgloss.JoinHorizontal(
		lipgloss.Top,
		mode,
		filepath,
		cursor,
	)

	return ui
}

func renderMode(m Model) string {
	var mode string
	text := ModeString[m.mode]

	switch m.mode {
	case NormalMode:
		mode = NormalModeStyle.Width(len(text) + 2).Render(text)
	case InsertMode:
		mode = InsertModeStyle.Width(len(text) + 2).Render(text)
	case CommandMode:
		mode = CommandModeStyle.Width(len(text) + 2).Render(text)
	case VisualMode:
		mode = VisualModelStyle.Width(len(text) + 2).Render(text)
	case SearchMode:
		mode = SearchModelStyle.Width(len(text) + 2).Render(text)
	default:
		mode = NormalModeStyle.Width(len(text) + 2).Render(text)
	}

	return mode
}

func renderFilePath(m Model) string {
	width_mode := len(ModeString[m.mode]) + 2
	width_cursor := len(strconv.Itoa(m.cursor.row)) + len(strconv.Itoa(m.cursor.column)) + 3
	width := m.width - width_mode - width_cursor

	return StatusLineStyle.Width(width).Render(" " + m.fileName)
}

func renderCursor(m Model) string {
	status := strings.Builder{}
	status.WriteString(strconv.Itoa(m.cursor.row))
	status.WriteString(":")
	status.WriteString(strconv.Itoa(m.cursor.column))

	var cursor string

	switch m.mode {
	case NormalMode:
		cursor = NormalModeStyle.Width(len(status.String()) + 2).Render(status.String())
	case InsertMode:
		cursor = InsertModeStyle.Width(len(status.String()) + 2).Render(status.String())
	case CommandMode:
		cursor = CommandModeStyle.Width(len(status.String()) + 2).Render(status.String())
	case VisualMode:
		cursor = VisualModelStyle.Width(len(status.String()) + 2).Render(status.String())
	case SearchMode:
		cursor = SearchModelStyle.Width(len(status.String()) + 2).Render(status.String())
	default:
		cursor = NormalModeStyle.Width(len(status.String()) + 2).Render(status.String())
	}

	return cursor
}

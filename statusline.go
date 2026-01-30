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
	case Normal:
		mode = NormalModeStyle.Width(len(text) + 2).Render(text)
	case Insert:
		mode = InsertModeStyle.Width(len(text) + 2).Render(text)
	case Command:
		mode = CommandModeStyle.Width(len(text) + 2).Render(text)
	case Visual:
		mode = VisualModelStyle.Width(len(text) + 2).Render(text)
	case Search:
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
	case Normal:
		cursor = NormalModeStyle.Width(len(status.String()) + 2).Render(status.String())
	case Insert:
		cursor = InsertModeStyle.Width(len(status.String()) + 2).Render(status.String())
	case Command:
		cursor = CommandModeStyle.Width(len(status.String()) + 2).Render(status.String())
	case Visual:
		cursor = VisualModelStyle.Width(len(status.String()) + 2).Render(status.String())
	case Search:
		cursor = SearchModelStyle.Width(len(status.String()) + 2).Render(status.String())
	default:
		cursor = NormalModeStyle.Width(len(status.String()) + 2).Render(status.String())
	}

	return cursor
}

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
		StatusLineStyle.Render("\t"),
		StatusLineStyle.Render(filepath),
		StatusLineStyle.Render("\t"),
		cursor,
	)

	return ui
}

func renderMode(m Model) string {
	var mode string

	switch m.mode {
	case Normal:
		text := "Normal"
		mode = NormalModeStyle.Width(len(text) + 2).Render(text)
	case Insert:
		text := "Insert"
		mode = InsertModeStyle.Width(len(text) + 2).Render(text)
	case Command:
		text := "Command"
		mode = CommandModeStyle.Width(len(text) + 2).Render(text)
	case Visual:
		text := "Visual"
		mode = VisualModelStyle.Width(len(text) + 2).Render(text)
	default:
		text := "Normal"
		mode = NormalModeStyle.Width(len(text) + 2).Render(text)
	}

	return mode
}

func renderFilePath(m Model) string {
	return m.fileName
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
	default:
		cursor = NormalModeStyle.Width(len(status.String()) + 2).Render(status.String())
	}

	return cursor
}

package main

import (
	"fmt"

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
		StatusLineStyle.Render(cursor),
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
	return fmt.Sprintf("%d:%d", m.cursor.row, m.cursor.column)
}

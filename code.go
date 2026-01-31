package main

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func RenderCode(m Model) string {
	code_pre := renderPreCode(m)
	code_selected := renderSelectedLine(m)
	code_post := renderPostCode(m)

	var ui string
	if code_pre == "" {
		ui = lipgloss.JoinVertical(
			lipgloss.Top,
			CursorStyle.Render(code_selected),
			CodeStyle.Render(code_post),
		)
	} else if code_post == "" {
		ui = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeStyle.Render(code_pre),
			CursorStyle.Render(code_selected),
		)
	} else {
		ui = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeStyle.Width(m.width).Render(code_pre),
			CursorStyle.Render(code_selected),
			CodeStyle.Width(m.width).Render(code_post),
		)
	}

	return ui
}

func renderPreCode(m Model) string {
	code := strings.Builder{}

	for index := 0; index < m.cursor.row; index++ {
		code.WriteString(m.buffer[index])
		if index != m.cursor.row-1 {
			code.WriteString("\n")
		}
	}

	return code.String()
}

func renderSelectedLine(m Model) string {
	code_pre := ""
	code_selected := ""
	code_post := ""

	if len(m.buffer) != 0 && len(m.buffer[m.cursor.row]) != 0 {
		line := m.buffer[m.cursor.row]
		index := m.cursor.column

		code_pre = line[:index]
		code_selected = string(line[min(index, max(0, len(line)-1))])
		code_post = line[min(index+1, len(line)):]
	} else {
		code_selected = " " // Insert space to show cursor
	}

	// Prevent cursor from widening on a tab
	if code_selected == "\t" {
		code_selected = " "
		code_post = "   " + code_post
	}

	ui := lipgloss.JoinHorizontal(
		lipgloss.Left,
		CodeStyle.Render(code_pre),
		CursorStyle.Render(code_selected),
		CodeStyle.Width(m.width).UnsetPaddingLeft().Render(code_post),
	)

	return ui
}

func renderPostCode(m Model) string {
	code := strings.Builder{}

	for index := min(m.cursor.row+1, len(m.buffer)); index < max(0, len(m.buffer)); index++ {
		code.WriteString(m.buffer[index])
		code.WriteString("\n")
	}

	// Fill window with empty space
	for i := 0; i < m.height-3-len(m.buffer); i++ {
		code.WriteString("\n")
	}

	return code.String()
}

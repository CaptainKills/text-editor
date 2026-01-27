package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderCode(m Model) string {
	code_pre := renderPreCode(m)
	code_selected := renderSelectedLine(m.buffer[m.cursor.row], m.cursor.column)
	code_post := renderPostCode(m)

	var ui string
	if code_pre == "" {
		ui = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeHighlightStyle.Render(code_selected),
			CodeStyle.Render(code_post),
		)
	} else if code_post == "" {
		ui = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeStyle.Render(code_pre),
			CodeHighlightStyle.Render(code_selected),
		)
	} else {
		ui = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeStyle.Render(code_pre),
			CodeHighlightStyle.Render(code_selected),
			CodeStyle.Render(code_post),
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

func renderSelectedLine(line string, index int) string {
	code_pre := ""
	code_selected := " "
	code_post := ""

	if len(line) != 0 {
		code_pre = line[:index]
		code_selected = string(line[min(index, max(0, len(line)-1))])
		code_post = line[min(index+1, len(line)):]
	}

	ui := lipgloss.JoinHorizontal(
		lipgloss.Left,
		CodeStyle.Render(code_pre),
		CodeHighlightStyle.Render(code_selected),
		CodeStyle.UnsetPaddingLeft().Render(code_post),
	)

	return ui
}

func renderPostCode(m Model) string {
	code := strings.Builder{}

	for index := min(m.cursor.row+1, len(m.buffer)); index < max(0, len(m.buffer)); index++ {
		code.WriteString(m.buffer[index])
		if index != len(m.buffer)-1 {
			code.WriteString("\n")
		}
	}

	return code.String()
}

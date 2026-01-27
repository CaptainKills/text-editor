package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderCode(m Model) string {
	code_pre := renderPreCode(m)
	code_highlight := renderSelectedLine(m.buffer[m.cursor.row])
	code_post := renderPostCode(m)

	var code string
	if code_pre == "" {
		code = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeHighlightStyle.Render(code_highlight),
			CodeStyle.Render(code_post),
		)
	} else if code_post == "" {
		code = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeStyle.Render(code_pre),
			CodeHighlightStyle.Render(code_highlight),
		)
	} else {
		code = lipgloss.JoinVertical(
			lipgloss.Top,
			CodeStyle.Render(code_pre),
			CodeHighlightStyle.Render(code_highlight),
			CodeStyle.Render(code_post),
		)
	}

	return code
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

func renderSelectedLine(line string) string {
	code := strings.Builder{}

	code.WriteString(line)

	return code.String()
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

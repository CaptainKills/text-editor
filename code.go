package main

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func RenderCode(m Model) string {
	builder := strings.Builder{}

	width := m.width - StatusColumnStyle.GetWidth()
	height := m.height - 3

	limit := min(m.cursor.offset+height, len(m.buffer))

	for index := m.cursor.offset; index < limit; index++ {
		builder.WriteString(m.buffer[index])
		if index != limit-1 {
			builder.WriteString("\n")
		}
	}

	code_layer := lipgloss.NewLayer(
		CodeStyle.Width(width).Height(height).Render(builder.String()),
	)

	highlight_layer := lipgloss.NewLayer(
		CursorStyle.Width(m.width).Render(m.buffer[m.cursor.row]),
	).Y(m.cursor.row - m.cursor.offset)

	layers := []*lipgloss.Layer{
		code_layer,
		highlight_layer,
	}

	compositor := lipgloss.NewCompositor(layers...)
	return compositor.Render()
}

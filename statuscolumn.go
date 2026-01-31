package main

import (
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
)

func RenderStatusColumn(m Model) string {
	builder := strings.Builder{}
	height := m.height - 3

	limit := min(m.cursor.offset+height, len(m.buffer))
	if len(m.buffer) != 0 {
		for index := m.cursor.offset; index < limit; index++ {
			builder.WriteString(strconv.Itoa(index))

			if index != limit-1 {
				builder.WriteString("\n")
			}
		}
	} else {
		builder.WriteString(strconv.Itoa(0))
	}

	rendered_numbers := StatusColumnStyle.Height(height).Render(builder.String())
	width := lipgloss.Width(rendered_numbers)

	number_layer := lipgloss.NewLayer(rendered_numbers)

	highlight_layer := lipgloss.NewLayer(
		HighlightStyle.Width(width).Render(strconv.Itoa(m.cursor.row)),
	).Y(m.cursor.row - m.cursor.offset)

	layers := []*lipgloss.Layer{
		number_layer,
		highlight_layer,
	}

	compositor := lipgloss.NewCompositor(layers...)
	return compositor.Render()
}

package main

import (
	"strconv"
	"strings"
)

func RenderStatusColumn(m Model) string {
	status := strings.Builder{}

	if len(m.buffer) != 0 {
		for index := range m.buffer {
			status.WriteString(strconv.Itoa(index))
			if index != len(m.buffer)-1 {
				status.WriteString("\n")
			}
		}
	} else {
		status.WriteString(strconv.Itoa(0))
	}

	height := m.height - StatusLineStyle.GetHeight() - CommandModeStyle.GetHeight() - 2

	return StatusColumnStyle.Height(height).Render(status.String())
}

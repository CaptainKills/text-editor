package main

import (
	"strconv"
	"strings"
)

func RenderNumberColumn(m Model) string {
	numbers := strings.Builder{}

	for index := range m.buffer {
		numbers.WriteString(strconv.Itoa(index))
		if index != len(m.buffer)-1 {
			numbers.WriteString("\n")
		}
	}

	return LineStyle.Render(numbers.String())
}

func RenderStatusLine(m Model) string {
	status := strings.Builder{}

	status.WriteString(m.fileName)
	status.WriteString(" - ")
	status.WriteString(strconv.Itoa(m.cursor.row))
	status.WriteString(":")
	status.WriteString(strconv.Itoa(m.cursor.column))

	return StatusStyle.Render(status.String())
}

func RenderCommandLine(m Model) string {
	return CommandStle.Render(m.command)
}

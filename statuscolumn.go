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

	return StatusColumnStyle.Render(status.String())
}

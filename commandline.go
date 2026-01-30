package main

func RenderCommandLine(m Model) string {
	return CommandStyle.Width(m.width).Render(m.command)
}

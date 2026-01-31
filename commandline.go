package main

func RenderCommandLine(m Model) string {
	return CommandLineStyle.Width(m.width).Render(m.command)
}

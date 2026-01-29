package main

func RenderCommandLine(m Model) string {
	return CommandStyle.Render(m.command)
}

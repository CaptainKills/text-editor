package main

import "charm.land/lipgloss/v2"

// ----- Code -----
var CodeStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#000000"))

var CursorStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#222222"))

// ----- Status Column -----
var StatusColumnStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#000000")).
	BorderStyle(lipgloss.NormalBorder()).
	BorderRight(true).
	AlignHorizontal(lipgloss.Right)

var HighlightStyle = StatusColumnStyle.
	Background(lipgloss.Color("#222222"))

// ----- Status Line -----
var StatusLineStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Left).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#222222"))

var NormalModeStyle = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Center).
	Bold(true).
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#3399FF"))

var InsertModeStyle = NormalModeStyle.
	Background(lipgloss.Color("#00CC99"))

var CommandModeStyle = NormalModeStyle.
	Background(lipgloss.Color("#FF9933"))

var VisualModelStyle = NormalModeStyle.
	Background(lipgloss.Color("#FF66FF"))

var SearchModelStyle = NormalModeStyle.
	Background(lipgloss.Color("#E6E600"))

// ----- Command Line -----
var CommandLineStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#222222"))

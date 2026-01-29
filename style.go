package main

import "github.com/charmbracelet/lipgloss"

// ----- Code -----
var CodeStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#000000"))

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#FFFFFF"))

// ----- Status Column -----
var StatusColumnStyle = lipgloss.NewStyle().
	// PaddingTop(1).
	PaddingLeft(2).
	BorderStyle(lipgloss.NormalBorder()).
	BorderRight(true).
	AlignHorizontal(lipgloss.Right)

// ----- Status Line -----
var StatusLineStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderTop(true)

var NormalModeStyle = lipgloss.NewStyle().
	Width(10).
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

// ----- Command Line -----
var CommandStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true)

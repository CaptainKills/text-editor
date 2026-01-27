package main

import "github.com/charmbracelet/lipgloss"

// ----- Code -----
var CodeStyle = lipgloss.NewStyle().
	PaddingLeft(2).
	Foreground(lipgloss.Color("#FFFFFF")).
	Background(lipgloss.Color("#000000"))

var CodeHighlightStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#000000")).
	Background(lipgloss.Color("#FFFFFF"))

// ----- Number Column -----
var LineStyle = lipgloss.NewStyle().
	// PaddingTop(1).
	PaddingLeft(2).
	BorderStyle(lipgloss.NormalBorder()).
	BorderRight(true)

// ----- Status Line -----
var StatusStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderTop(true)

// ----- Command Line -----
var CommandStle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottom(true)

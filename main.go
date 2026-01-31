package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
)

func main() {
	fileName := "main.go"
	path, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatalf("Could not create filepath: %v\n", err)
	}

	lines, err := ReadFile(path)
	if err != nil {
		log.Fatalf("Could not read file: %v\n", err)
	}

	f, err := tea.LogToFile("debug.log", "")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	m := Model{fileName: path, buffer: lines, log: f}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("There's been an error: %v", err)
	}
}

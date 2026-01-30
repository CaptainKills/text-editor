package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func readFile(fileName string) ([]string, error) {
	inputFile, err := os.Open(fileName)
	if err != nil {
		return []string{}, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func main() {
	fileName := "main.go"
	path, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatalf("Could not create filepath: %v\n", err)
	}

	lines, err := readFile(path)
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

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("There's been an error: %v", err)
	}
}

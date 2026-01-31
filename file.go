package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
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
		line := strings.ReplaceAll(scanner.Text(), "\t", "    ")
		lines = append(lines, line)
	}

	return lines, nil
}

func WriteFile(fileName string, buffer []string) {
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file! %q\n", err)
	}
	defer outputFile.Close()

	for index := range buffer {
		line := strings.ReplaceAll(buffer[index], "    ", "\t")
		outputFile.WriteString(line)

		if index != len(buffer)-1 {
			outputFile.WriteString("\n")
		}
	}
}

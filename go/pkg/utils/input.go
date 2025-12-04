package utils

import (
	"bufio"
	"os"
	"strings"
)

// ReadInput reads the entire file and returns it as a string
func ReadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads the file and returns it as a slice of lines
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// MustReadInput reads the file and panics on error
func MustReadInput(filename string) string {
	input, err := ReadInput(filename)
	if err != nil {
		panic(err)
	}
	return input
}

// MustReadLines reads lines from file and panics on error
func MustReadLines(filename string) []string {
	lines, err := ReadLines(filename)
	if err != nil {
		panic(err)
	}
	return lines
}

// JoinLines joins lines with newline separator
func JoinLines(lines []string) string {
	return strings.Join(lines, "\n")
}

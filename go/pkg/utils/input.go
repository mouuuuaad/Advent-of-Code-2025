package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

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

func MustReadInput(filename string) string {
	input, err := ReadInput(filename)
	if err != nil {
		panic(err)
	}
	return input
}

func MustReadLines(filename string) []string {
	lines, err := ReadLines(filename)
	if err != nil {
		panic(err)
	}
	return lines
}

func JoinLines(lines []string) string {
	return strings.Join(lines, "\n")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")
	dialPos := 50
	zeroCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		direction := line[0]
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		if direction == 'R' {
			dialPos = (dialPos + distance) % 100
		} else if direction == 'L' {
			dialPos = (dialPos - distance) % 100
			if dialPos < 0 {
				dialPos += 100
			}
		}

		if dialPos == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")
	dialPos := 50
	zeroCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		direction := line[0]
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		step := 0
		if direction == 'R' {
			step = 1
		} else {
			step = -1
		}

		for i := 0; i < distance; i++ {
			dialPos = (dialPos + step) % 100
			if dialPos < 0 {
				dialPos += 100
			}
			if dialPos == 0 {
				zeroCount++
			}
		}
	}

	return zeroCount
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	input := strings.Join(lines, "\n")
	fmt.Println("Part 1 Password:", solvePart1(input))
	fmt.Println("Part 2 Password:", solvePart2(input))
}

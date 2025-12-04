package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")
	grid := [][]rune{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	accessible := 0

	dirs := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '@' {
				neighbors := 0
				for _, d := range dirs {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == '@' {
						neighbors++
					}
				}
				if neighbors < 4 {
					accessible++
				}
			}
		}
	}

	return accessible
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")
	grid := [][]rune{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	totalRemoved := 0

	dirs := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for {
		accessible := [][]int{}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] == '@' {
					neighbors := 0
					for _, d := range dirs {
						ni, nj := i+d[0], j+d[1]
						if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == '@' {
							neighbors++
						}
					}
					if neighbors < 4 {
						accessible = append(accessible, []int{i, j})
					}
				}
			}
		}

		if len(accessible) == 0 {
			break
		}
		for _, pos := range accessible {
			grid[pos[0]][pos[1]] = '.'
		}
		totalRemoved += len(accessible)
	}

	return totalRemoved
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
	fmt.Println("Part 1 Accessible Rolls:", solvePart1(input))
	fmt.Println("Part 2 Total Removed:", solvePart2(input))
}
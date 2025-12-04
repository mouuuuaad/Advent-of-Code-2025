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
	totalJoltage := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		maxJoltage := 0
		digits := []int{}
		for _, r := range line {
			d, _ := strconv.Atoi(string(r))
			digits = append(digits, d)
		}

		for i := 0; i < len(digits); i++ {
			for j := i + 1; j < len(digits); j++ {
				joltage := digits[i]*10 + digits[j]
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}
		totalJoltage += maxJoltage
	}

	return totalJoltage
}

func solvePart2(input string) int64 {
	lines := strings.Split(input, "\n")
	var totalJoltage int64 = 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		digits := []int{}
		for _, r := range line {
			d, _ := strconv.Atoi(string(r))
			digits = append(digits, d)
		}

		k := 12
		stack := []int{}
		n := len(digits)

		for i, digit := range digits {
			for len(stack) > 0 && stack[len(stack)-1] < digit && len(stack)-1+(n-i) >= k {
				stack = stack[:len(stack)-1]
			}
			if len(stack) < k {
				stack = append(stack, digit)
			}
		}

		var val int64 = 0
		for _, d := range stack {
			val = val*10 + int64(d)
		}
		totalJoltage += val
	}

	return totalJoltage
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
	fmt.Println("Part 1 Total Output Joltage:", solvePart1(input))
	fmt.Println("Part 2 Total Output Joltage:", solvePart2(input))
}

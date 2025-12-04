package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidPart1(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

func solvePart1(input string) int {
	ranges := strings.Split(input, ",")
	totalSum := 0

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidPart1(i) {
				totalSum += i
			}
		}
	}
	return totalSum
}

func isInvalidPart2(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)
	for subLen := 1; subLen <= length/2; subLen++ {
		if length%subLen == 0 {
			pattern := s[:subLen]
			repeated := strings.Repeat(pattern, length/subLen)
			if s == repeated {
				return true
			}
		}
	}
	return false
}
func solvePart2(input string) int {
	ranges := strings.Split(input, ",")
	totalSum := 0

	for _, r := range ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			if isInvalidPart2(i) {
				totalSum += i
			}
		}
	}
	return totalSum
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

	input := strings.Join(lines, "")
	fmt.Println("Part 1 Sum of invalid IDs:", solvePart1(input))
	fmt.Println("Part 2 Sum of invalid IDs:", solvePart2(input))
}

# Day 1: Safe Dial Puzzle

**Problem**: [Advent of Code 2025 - Day 1](https://adventofcode.com/2025/day/1)

## Problem Description

Simulate a safe dial starting at position 50. Follow instructions to rotate left (L) or right (R) by specified distances. The dial has 100 positions (0-99) and wraps around.

**Part 1**: Count how many times the dial crosses position 0 after completing all moves.

**Part 2**: Count how many times the dial crosses position 0 during the moves (not just at the end of each instruction).

## Solution Approach

**Part 1**:
- Start at position 50
- Process each instruction, update position with modulo arithmetic
- Check if position equals 0 after each instruction
- Count occurrences

**Part 2**:
- Same as Part 1, but step through each position change one at a time
- For each step in a multi-position move, check if we cross 0
- Provides more granular tracking

## How to Run

```bash
# From go/day01 directory
go run main.go

# Run tests
go test -v

# Using Makefile from project root
make go-run DAY=1
```

## Complexity

- **Time**: O(n) where n is the number of instructions for Part 1, O(n*m) for Part 2 where m is average distance
- **Space**: O(1)
# Day 2: Invalid Product IDs

**Problem**: [Advent of Code 2025 - Day 2](https://adventofcode.com/2025/day/2)

## Problem Description

Identify invalid product IDs within given ranges based on repeating digit patterns.

**Part 1**: Find IDs where the number splits perfectly in half with identical halves (e.g., 1212, 567567).

**Part 2**: Find IDs where the entire number is made of a repeating pattern of any length (e.g., 111, 123123, 4545).
## Solution Approach

**Part 1**:
- Check if number length is even
- Split into two halves
- Compare if halves are identical

**Part 2**:
- Try all possible pattern lengths from 1 to length/2
- For each length, check if the number is that pattern repeated
- Return true if any pattern works

Both parts iterate through the given ranges and sum up all invalid IDs.

## How to Run

```bash
# From go/day02 directory
go run main.go

# Run tests
go test -v

# Using Makefile from project root
make go-run DAY=2
```

## Complexity

- **Time**: O(n * m) where n is range size, m is average number length
- **Space**: O(1)

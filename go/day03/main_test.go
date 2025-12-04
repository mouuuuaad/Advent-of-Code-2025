package main

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	expected := 357
	result := solvePart1(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `987654321111111
811111111111119
234234234234278
818181911112111`
	var expected int64 = 3121910778619
	result := solvePart2(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

package main

import "testing"

func TestSolvePart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := 3
	result := solvePart1(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := 6
	result := solvePart2(input)

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
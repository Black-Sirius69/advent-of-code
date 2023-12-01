package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	if result := Part1(input); result != 142 {
		t.Errorf("Part1 FAILED. Expected 142 got %d", result)
	} else {
		t.Logf("Part1 PASSED.")
	}
}

func TestPart2(t *testing.T) {
	input := `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
    `
	if result := Part2(input); result != 281 {
		t.Errorf("Part2 FAILED. Expected 142 got %d", result)
	} else {
		t.Logf("Part2 PASSED.")
	}
}

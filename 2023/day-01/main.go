package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func checkDigit(slice string) int64 {
	if strings.HasPrefix(slice, "one") {
		return 1
	}
	if strings.HasPrefix(slice, "two") {
		return 2
	}
	if strings.HasPrefix(slice, "three") {
		return 3
	}
	if strings.HasPrefix(slice, "four") {
		return 4
	}
	if strings.HasPrefix(slice, "five") {
		return 5
	}
	if strings.HasPrefix(slice, "six") {
		return 6
	}
	if strings.HasPrefix(slice, "seven") {
		return 7
	}
	if strings.HasPrefix(slice, "eight") {
		return 8
	}
	if strings.HasPrefix(slice, "nine") {
		return 9
	}

	return 0
}

func Part1(input string) int64 {
	var sum int64 = 0

	for _, line := range strings.Fields(input) {
		var left int64 = -1
		var right int64 = -1

		for _, char := range line {
			if unicode.IsDigit(rune(char)) {
				if left == -1 {
					left = int64(char - '0')
				} else {
					right = int64(char - '0')
				}
			}
		}

		if right == -1 {
			right = left
		}

		sum += left*10 + right
	}

	return sum
}

func Part2(input string) int64 {
	var sum int64 = 0

	for _, line := range strings.Fields(input) {
		var left int64 = -1
		var right int64 = -1

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if left == -1 {
					left = int64(line[i] - '0')
				} else {
					right = int64(line[i] - '0')
				}
			}

			if val := checkDigit(line[i:]); val != 0 {
				if left == -1 {
					left = val
				} else {
					right = val
				}
			}

		}

		if right == -1 {
			right = left
		}
		sum += left*10 + right
	}

	return sum
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(Part1(string(input)))
	fmt.Println(Part2(string(input)))
}

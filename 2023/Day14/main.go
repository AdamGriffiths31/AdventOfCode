package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	fmt.Printf("Lines: %v\n", lines)
	grid := rotate(lines)

	fmt.Printf("Grid: %v\n", grid)
	newGrid := solve(grid)

	rotated := rotate(newGrid)
	totalScore := 0
	for i, line := range rotated {
		occurences := countOccurences(line, "O")
		score := len(line) - i
		totalScore += occurences * score
	}
	fmt.Printf("Total score: %v\n", totalScore)
}

func solve(grid []string) []string {
	newGrid := make([]string, len(grid))
	for _, line := range grid {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "#")
		row := ""
		for _, part := range parts {
			row += sort(part) + "#"
		}
		newGrid = append(newGrid, row)
	}
	return newGrid[len(grid):]
}

func countOccurences(input, char string) int {
	count := 0
	for _, c := range input {
		if string(c) == char {
			count++
		}
	}
	return count
}

func sort(input string) string {
	chars := strings.Split(input, "")
	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] < chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return strings.Join(chars, "")
}

func rotate(input []string) []string {
	output := make([]string, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		for _, str := range input {
			if i < len(str) {
				output[i] += string(str[i])
			}
		}
	}
	return output
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	blocks := strings.Split(string(input), "\n\n")

	total := 0
	for _, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		row := solve(lines)
		total += row * 100
		col := solve(rotate(lines))
		total += col
	}
	fmt.Println(total)
}

func solve(input []string) int {
	for i := 1; i < len(input); i++ {
		above := make([]string, len(input[:i]))
		below := make([]string, len(input[i:]))
		above = reverse(input[:i])
		below = input[i:]

		newAboveLen := len(below)
		if len(above) < len(below) {
			newAboveLen = len(above)
		}
		newBelowLen := len(above)
		if len(below) < len(above) {
			newBelowLen = len(below)
		}

		above = above[:newAboveLen]
		below = below[:newBelowLen]

		if countDifferences(above, below) == 1 {
			return i
		}
	}
	return 0
}

func reverse(input []string) []string {
	output := make([]string, len(input))
	for i, v := range input {
		output[len(input)-1-i] = v
	}
	return output
}

func areEqual(input1 []string, input2 []string) bool {
	if len(input1) != len(input2) {
		return false
	}

	for i, v := range input1 {
		if v != input2[i] {
			return false
		}
	}
	return true
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

func countDifferences(above, below []string) int {
	count := 0
	for i := range above {
		for j := range above[i] {
			if above[i][j] != below[i][j] {
				count++
			}
		}
	}
	return count
}

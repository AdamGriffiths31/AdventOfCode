package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var seenGrids = make(map[string]bool)
var array = []string{}
var grid = []string{}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	grid := strings.Split(string(input), "\n")

	found := false
	count := 0
	for found == false {
		count++
		grid = solve(grid)
		fmt.Println(count)
		printGrid(grid)
		if seenGrids[strings.Join(grid, "")] == true {
			found = true
			println(count)
		}
		seenGrids[strings.Join(grid, "")] = true
		array = append(array, strings.Join(grid, ""))
	}

	first := 0
	for i, str := range array {
		if str == strings.Join(grid, "") {
			first = 1 + i
			break
		}
	}

	res := array[(1000000000-first)%(count-first)+first-1]
	grid = splitString(res, len(grid))
	sum := 0
	for i, row := range grid {
		if row == "" {
			continue
		}
		occurences := countOccurences(row, "O")
		fmt.Println(occurences, len(grid)-i)
		sum += occurences * (len(grid) - i)
	}

	fmt.Println(sum)
}

func solve(grid []string) []string {
	for i := 0; i < 4; i++ {
		newGrid := []string{}
		grid = rotate(grid)
		for _, line := range grid {
			parts := strings.Split(line, "#")
			sortedParts := make([]string, len(parts))
			for j, part := range parts {
				sortedParts[j] = sort(part)
			}

			newGrid = append(newGrid, reverse(strings.Join(sortedParts, "#")))
		}
		grid = newGrid
	}
	return grid
}

func printGrid(grid []string) {
	for _, line := range grid {
		fmt.Println(line)
	}
	fmt.Println()
}

func reverse(input string) string {
	output := ""
	for _, c := range input {
		output = string(c) + output
	}
	return output
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

func splitString(input string, chunkSize int) []string {
	var result []string

	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}

	return result
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

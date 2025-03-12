package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	fmt.Printf("Sum: %d\n", sum)

	actionRegex := regexp.MustCompile(`do\(\)|don't\(\)`)
	actions := actionRegex.FindAllStringIndex(input, -1)
	enabled := true
	var result int
	nextStartIndex := 0

	for _, action := range actions {
		start, end := action[0], action[1]
		if enabled {
			segment := input[nextStartIndex:start]
			mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
			matches := mulRegex.FindAllStringSubmatch(segment, -1)
			for _, match := range matches {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				result += x * y
			}
		}

		actionText := input[start:end]
		switch actionText {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		}

		nextStartIndex = end
	}

	if enabled {
		segment := input[nextStartIndex:]
		mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := mulRegex.FindAllStringSubmatch(segment, -1)
		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result += x * y
		}
	}

	fmt.Printf("Result: %d\n", result)
}

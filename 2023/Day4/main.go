package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	part1 := 0
	row := 0
	scanner := bufio.NewScanner(file)
	cards := map[int]int{}
	for scanner.Scan() {
		row++
		part1 += solve(row, scanner.Text(), cards)
	}
	part2 := 0
	for _, v := range cards {
		part2 += v
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func solve(gameNumber int, input string, cards map[int]int) int {
	points := 0
	winningNumbers := make([]string, 0)
	parts := strings.Split(input, "|")
	winningParts := strings.Split(parts[0], ":")[1:]

	for _, part := range winningParts {
		numbers := strings.Split(part, " ")
		for _, number := range numbers {
			if number == "" {
				continue
			}
			winningNumbers = append(winningNumbers, number)
		}
	}

	winners := 0
	for _, part := range parts[1:] {
		numbers := strings.Split(part, " ")
		for _, number := range numbers {
			if number == "" {
				continue
			}
			if contains(winningNumbers, number) {
				winners++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
	}
	cards[gameNumber]++
	for j := 1; j <= winners; j++ {
		cards[gameNumber+j] += cards[gameNumber]
	}
	return points
}

func contains(numbers []string, number string) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

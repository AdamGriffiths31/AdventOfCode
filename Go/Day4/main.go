package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(input)

	var score int

	for scanner.Scan() {
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)

		if startSecond >= startFirst && endSecond <= endFirst || startFirst >= startSecond && endFirst <= endSecond {
			score++
		}
	}
	fmt.Println(score)
	partB()
}

func partB() {
	input, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(input)

	var score int

	for scanner.Scan() {
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)
		if startSecond <= endFirst && endSecond >= startFirst || startFirst <= endSecond && endFirst >= startSecond {
			score++
		}
	}
	fmt.Println(score)
}

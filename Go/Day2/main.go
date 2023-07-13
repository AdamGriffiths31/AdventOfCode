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
	scores := map[string]int{"B X": 1, "C Y": 2, "A Z": 3, "A X": 4, "B Y": 5, "C Z": 6, "C X": 7, "A Y": 8, "B Z": 9}

	for scanner.Scan() {
		score += scores[scanner.Text()]
	}
	fmt.Println(score)
	part2()
}

func part2() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	var score int
	scores := map[string]int{"B X": 1, "C X": 2, "A X": 3, "A Y": 4, "B Y": 5, "C Y": 6, "C Z": 7, "A Z": 8, "B Z": 9}

	for scanner.Scan() {
		score += scores[scanner.Text()]
	}
	fmt.Println(score)
}

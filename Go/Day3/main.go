package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")

	values := make(map[rune]int)
	counter := 1
	for i := 'a'; i <= 'z'; i++ {
		values[i] = counter
		counter++
	}
	for i := 'A'; i <= 'Z'; i++ {
		values[i] = counter
		counter++
	}

	var score int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		divide := length / 2

		hashMap := make(map[rune]int)
		for _, char := range line[:divide] {
			hashMap[char]++
		}

		for _, char := range line[divide:] {
			if _, exists := hashMap[char]; exists {
				fmt.Println(string(char))
				fmt.Println("score", values[char])
				score += values[char]
				break
			}
		}
	}

	fmt.Println(score)
}

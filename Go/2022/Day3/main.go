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
				score += values[char]
				break
			}
		}
	}

	fmt.Println(score)
	partB()
}

func partB() {
	input, _ := os.Open("input.txt")
	var score int

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

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row1 := creatMaps(scanner.Text())
		scanner.Scan()
		row2 := creatMaps(scanner.Text())
		scanner.Scan()
		row3 := creatMaps(scanner.Text())

		for elf1Item := range row1 {
			if row2[elf1Item] && row3[elf1Item] {
				score += values[elf1Item]
				break
			}
		}

	}
	fmt.Println(score)
}

func creatMaps(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, item := range items {
		set[item] = true
	}
	return
}

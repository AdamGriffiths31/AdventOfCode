package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	parts := strings.Split(string(input), ",")
	total := 0
	for _, part := range parts {
		if part == "" || part == "\n" {
			continue
		}
		val := solve(part)
		fmt.Printf("%s %d\n", part, val)
		total += val
	}
	fmt.Printf("Total: %d\n", total)
}

func solve(input string) int {
	start := 0
	for _, char := range input {
		if char == '\n' {
			continue
		}
		start += int(char)
		start *= 17
		start %= 256
	}
	return start
}

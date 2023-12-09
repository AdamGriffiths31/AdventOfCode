package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	parts := strings.Split(string(input), "\n")

	total := 0
	for _, part := range parts {
		if part == "" {
			continue
		}
		total += extrapolate(parse(part))
	}
	fmt.Println(total)
}

func parse(input string) []int {
	parts := strings.Split(input, " ")
	output := make([]int, len(parts))
	for i, part := range parts {
		output[i], _ = strconv.Atoi(part)
	}
	return output
}

func extrapolate(input []int) int {
	if allZero(input) {
		return 0
	}

	deltas := make([]int, len(input)-1)
	for i := 0; i < len(input)-1; i++ {
		deltas[i] = input[i+1] - input[i]
	}

	diff := extrapolate(deltas)
	return input[0] - diff
}

func allZero(input []int) bool {
	for _, v := range input {
		if v != 0 {
			return false
		}
	}
	return true
}

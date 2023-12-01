package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxTotal := 0
	current := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > maxTotal {
				maxTotal = current
			}
			current = 0
		}
		value, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		current += int(value)
	}

	fmt.Println(maxTotal)
	part2()
}

func part2() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxTotal1 := 0
	maxTotal2 := 0
	maxTotal3 := 0
	current := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if current > maxTotal3 {
				maxTotal3 = current
			}
			if maxTotal3 > maxTotal2 {
				maxTotal3, maxTotal2 = maxTotal2, maxTotal3
			}
			if maxTotal2 > maxTotal1 {
				maxTotal2, maxTotal1 = maxTotal1, maxTotal2
			}
			current = 0
		}
		value, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		current += int(value)
	}

	fmt.Println(maxTotal1 + maxTotal2 + maxTotal3)
}

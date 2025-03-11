package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	result := 0
	fixableCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Fields(line)
		var numbers []int

		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}

		if isSafe(numbers) {
			result++
		}

		if !isSafe(numbers) && fixList(numbers) {
			fixableCount++
		}
	}

	fmt.Println("Result:", result)
	fmt.Println("Fixable Count:", fixableCount)
}

func isSafe(numbers []int) bool {
	isIncreasing := numbers[0] < numbers[1]

	for i := 1; i < len(numbers); i++ {
		diff := math.Abs(float64(numbers[i] - numbers[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}

		if isIncreasing && numbers[i] < numbers[i-1] {
			return false
		}
		if !isIncreasing && numbers[i] > numbers[i-1] {
			return false
		}
	}

	return true
}

func fixList(numbers []int) bool {
	for i := range numbers {
		modified := append([]int{}, numbers[:i]...)
		modified = append(modified, numbers[i+1:]...)

		if isSafe(modified) {
			return true
		}
	}
	return false
}

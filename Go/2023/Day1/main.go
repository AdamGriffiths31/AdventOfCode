package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	result := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += getLineValue(scanner.Text())
	}

	fmt.Printf("Result: %d\n", result)
	partTwo()
}

func partTwo() {
	result := 0

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += getLineValuePartTwo(scanner.Text())
	}

	fmt.Printf("Result: %d\n", result)
}

func getLineValuePartTwo(line string) int {
	first := ""
	last := ""

	for i, char := range line {
		if isInt(string(char)) {
			if first == "" {
				first = string(char)
				last = string(char)
			} else {
				last = string(char)
			}
		} else {
			for _, num := range numbers {
				if i+len(num) > len(line) {
					continue
				}
				possible := string(line[i : i+len(num)])
				if possible == num {
					if first == "" {
						first = strconv.Itoa(intReturn(num))
						last = strconv.Itoa(intReturn(num))
					} else {
						last = strconv.Itoa(intReturn(num))
					}

				}
			}
		}
	}

	fmt.Printf("First: %s, Last: %s\n", first, last)
	returnval, _ := strconv.Atoi(first + last)
	return returnval
}

func getLineValue(line string) int {
	start := 0
	end := len(line) - 1
	result := ""

	for start <= end {
		if isInt(string(line[start])) {
			result = string(line[start])
			break
		}
		start++
	}

	for start <= end {
		if isInt(string(line[end])) {
			result += string(line[end])
			break
		}
		end--
	}

	returnValue, _ := strconv.Atoi(result)
	return returnValue
}

func isInt(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func intReturn(st string) int {
	switch st {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		d, _ := strconv.Atoi(st)
		return d
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	parts := strings.Split(string(input), ",")
	boxes := make([][]string, 256)
	focalLengths := make(map[string]int)

	for _, part := range parts {
		if part == "" || part == "\n" {
			continue
		}
		if strings.Contains(part, "-") {
			label := part[:strings.Index(part, "-")]
			idx := solve(label)
			if contains(boxes[idx], label) {
				boxes[idx] = remove(boxes[idx], label)
			}
		} else {
			parts := strings.Split(part, "=")
			label := parts[0]
			length, _ := strconv.Atoi(parts[1])
			idx := solve(label)
			if !contains(boxes[idx], label) {
				boxes[idx] = append(boxes[idx], label)
			}
			focalLengths[label] = length
		}
	}
	total := 0
	for boxNumber, box := range boxes {
		for lensSlot, label := range box {
			fmt.Printf("%d %d %d\n", boxNumber+1, lensSlot+1, focalLengths[label])
			total += (boxNumber + 1) * (lensSlot + 1) * focalLengths[label]
		}
	}
	fmt.Println(total)
}

func remove(slice []string, str string) []string {
	for i, s := range slice {
		if s == str {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
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

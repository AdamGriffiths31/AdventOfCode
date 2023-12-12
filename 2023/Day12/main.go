package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		config := parts[0]
		counts := parts[1]

		count := convertToArray(counts)
		res := solve(config, count)
		total += res

		fmt.Printf("[%s] %+v   (%d)\n", config, count, res)
	}
	fmt.Println(total)
}

func convertToArray(nums string) []int {
	var arr []int
	for _, num := range strings.Split(nums, ",") {
		num = strings.TrimSpace(num)
		n, _ := strconv.Atoi(num)
		arr = append(arr, n)
	}
	return arr
}

func solve(config string, count []int) int {
	if config == "" {
		if len(count) == 0 {
			return 1
		}
		return 0
	}
	if len(count) == 0 {
		if strings.Contains(config, "#") {
			return 0
		}
		return 1
	}
	result := 0

	if config[0] == '.' || config[0] == '?' {
		result += solve(config[1:], count)
	}

	if config[0] == '#' || config[0] == '?' {
		if count[0] <= len(config) && !strings.Contains(config[:count[0]], ".") && (count[0] == len(config) || config[count[0]] != '#') {
			size := count[0] + 1
			if size >= len(config) {
				result += solve(config[len(config):], count[1:])
			} else {
				result += solve(config[count[0]+1:], count[1:])
			}
		}
	}
	return result
}

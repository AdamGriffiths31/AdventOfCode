package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var left, right []int
	leftCount := make(map[int]int)

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])

		left = append(left, l)
		right = append(right, r)

		leftCount[l]++
	}

	sort.Ints(left)
	sort.Ints(right)

	result := 0
	for i := 0; i < len(left); i++ {
		result += int(math.Abs(float64(left[i] - right[i])))
		println(left[i], right[i])
	}

	totalSum := 0
	for _, r := range right {
		totalSum += r * leftCount[r]
	}

	fmt.Println("Result:", result)
	fmt.Println("Total Sum:", totalSum)
}

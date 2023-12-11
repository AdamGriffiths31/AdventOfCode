package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	lines, _ := ioutil.ReadFile("input.txt")
	parts := strings.Split(string(lines), "\n")
	var grid []string
	for _, line := range parts {
		grid = append(grid, line)
	}

	var emptyRows []int
	for r, row := range grid {
		if allDots(row) {
			emptyRows = append(emptyRows, r)
		}
	}

	var emptyCols []int
	for c, col := range zip(grid) {
		if allDots(col) {
			emptyCols = append(emptyCols, c)
		}
	}

	var points [][]int
	for r, row := range grid {
		for c, ch := range row {
			if ch == '#' {
				points = append(points, []int{r, c})
			}
		}
	}

	// Print results
	fmt.Println("Empty Rows:", emptyRows)
	fmt.Println("Empty Columns:", emptyCols)
	fmt.Println("Points with '#':", points)

	total := 0
	totalPart2 := 0
	scale := 2
	scalePart2 := 1000000
	for i, point := range points {
		r1, c1 := point[0], point[1]
		for _, point2 := range points[:i] {
			r2, c2 := point2[0], point2[1]
			for row := min(r1, r2); row < max(r1, r2); row++ {
				if contains(emptyRows, row) {
					total += scale
					totalPart2 += scalePart2
				} else {
					total++
				}
			}

			for c := min(c1, c2); c < max(c1, c2); c++ {
				if contains(emptyCols, c) {
					total += scale
					totalPart2 += scalePart2
				} else {
					total++
				}
			}
		}
	}
	fmt.Println("Total:", total)
	fmt.Println("Total Part 2:", totalPart2)
}

func allDots(s string) bool {
	for _, ch := range s {
		if ch != '.' {
			return false
		}
	}
	return true
}

func zip(grid []string) []string {
	transposed := make([]string, len(grid[0]))
	for i := range transposed {
		for _, row := range grid {
			if i < len(row) {
				transposed[i] += string(row[i])
			}
		}
	}
	return transposed
}

func contains(s []int, n int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

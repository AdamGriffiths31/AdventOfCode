package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type point struct {
	r, c, dr, dc int
}

type cord struct {
	r, c int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	parts := strings.Split(string(input), "\n")

	grid := make([]string, len(parts))
	for i, part := range parts {
		grid[i] = part
	}

	highest := 0
	for row := range grid {
		highest = max(highest, solve(grid, row, -1, 0, 1))
		highest = max(highest, solve(grid, row, len(grid[0]), 0, -1))
	}

	for col := range grid[0] {
		highest = max(highest, solve(grid, -1, col, 1, 0))
		highest = max(highest, solve(grid, len(grid), col, -1, 0))
	}

	fmt.Println(highest)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(grid []string, r, c, dr, dc int) int {
	start := point{r, c, dr, dc}
	seen := make(map[point]bool)
	queue := []point{start}

	for len(queue) > 0 {
		poppedPoint := queue[0]
		queue = queue[1:]
		r, c, dr, dc := poppedPoint.r, poppedPoint.c, poppedPoint.dr, poppedPoint.dc
		r += dr
		c += dc

		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			continue
		}

		char := grid[r][c]
		if char == '.' || (char == '-' && dc != 0) || (char == '|' && dr != 0) {
			if !seen[point{r, c, dr, dc}] {
				seen[point{r, c, dr, dc}] = true
				queue = append(queue, point{r, c, dr, dc})
			}
		} else if char == '/' {
			dr, dc = -dc, -dr
			if !seen[point{r, c, dr, dc}] {
				seen[point{r, c, dr, dc}] = true
				queue = append(queue, point{r, c, dr, dc})
			}
		} else if char == '\\' {
			dr, dc = dc, dr
			if !seen[point{r, c, dr, dc}] {
				seen[point{r, c, dr, dc}] = true
				queue = append(queue, point{r, c, dr, dc})
			}
		} else {
			var directons [][]int
			if char == '|' {
				directons = [][]int{{1, 0}, {-1, 0}}
			} else {
				directons = [][]int{{0, 1}, {0, -1}}
			}
			for _, direction := range directons {
				dr, dc = direction[0], direction[1]
				if !seen[point{r, c, dr, dc}] {
					seen[point{r, c, dr, dc}] = true
					queue = append(queue, point{r, c, dr, dc})
				}
			}
		}
	}

	cords := map[cord]struct{}{}
	for point := range seen {
		cords[cord{point.r, point.c}] = struct{}{}
	}
	return len(cords)
}

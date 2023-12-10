package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	var sr, sc int
	seen := make(map[[2]int]bool)
	var q = make([][2]int, 0)

	for r, row := range grid {
		for c, ch := range row {
			if ch == 'S' {
				sr, sc = r, c
				break
			}
		}
	}

	seen[[2]int{sr, sc}] = true
	q = append(q, [2]int{sr, sc})

	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		ch := grid[r][c]

		if r > 0 && (ch == 'S' || ch == '|' || ch == 'J' || ch == 'L') && (grid[r-1][c] == '|' || grid[r-1][c] == '7' || grid[r-1][c] == 'F') && !seen[[2]int{r - 1, c}] {
			seen[[2]int{r - 1, c}] = true
			q = append(q, [2]int{r - 1, c})
		}

		if r < len(grid)-1 && (ch == 'S' || ch == '|' || ch == '7' || ch == 'F') && (grid[r+1][c] == '|' || grid[r+1][c] == 'J' || grid[r+1][c] == 'L') && !seen[[2]int{r + 1, c}] {
			seen[[2]int{r + 1, c}] = true
			q = append(q, [2]int{r + 1, c})
		}

		if c > 0 && (ch == 'S' || ch == '-' || ch == 'J' || ch == '7') && (grid[r][c-1] == '-' || grid[r][c-1] == 'L' || grid[r][c-1] == 'F') && !seen[[2]int{r, c - 1}] {
			seen[[2]int{r, c - 1}] = true
			q = append(q, [2]int{r, c - 1})
		}

		if c < len(grid[r])-1 && (ch == 'S' || ch == '-' || ch == 'L' || ch == 'F') && (grid[r][c+1] == '-' || grid[r][c+1] == 'J' || grid[r][c+1] == '7') && !seen[[2]int{r, c + 1}] {
			seen[[2]int{r, c + 1}] = true
			q = append(q, [2]int{r, c + 1})
		}
	}

	fmt.Printf("seen %v\n", len(seen)/2)
}


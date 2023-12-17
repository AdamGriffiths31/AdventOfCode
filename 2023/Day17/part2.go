package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	heat   int
	row    int
	col    int
	rowDir int
	colDir int
	count  int
}

type direction struct {
	row int
	col int
}

var directions = []direction{
	{row: 0, col: 1},
	{row: 1, col: 0},
	{row: 0, col: -1},
	{row: -1, col: 0},
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	lines := strings.Split(string(input), "\n")

	var grid = [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	fmt.Println("Part 2:")
	solve(grid)
}

func solve(grid [][]int) {

	seen := map[string]bool{}
	start := point{heat: 0, row: 0, col: 0, rowDir: 0, colDir: 0, count: 0}
	queue := []point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.row == len(grid)-1 && current.col == len(grid[0])-1 && current.count >= 4 {
			fmt.Println("Total:", current.heat)
			return
		}

		key := fmt.Sprintf("%d,%d,%d,%d,%d", current.row, current.col, current.rowDir, current.colDir, current.count)
		if seen[key] {
			continue
		}

		seen[key] = true

		if current.count < 10 && (current.rowDir != 0 || current.colDir != 0) {
			nextRow := current.row + current.rowDir
			nextCol := current.col + current.colDir
			if 0 <= nextRow && nextRow < len(grid) && 0 <= nextCol && nextCol < len(grid[0]) {
				next := point{heat: current.heat + grid[nextRow][nextCol], row: nextRow, col: nextCol, rowDir: current.rowDir, colDir: current.colDir, count: current.count + 1}
				queue = addToQueue(queue, next)
			}
		}

		if current.count >= 4 || (current.rowDir == 0 && current.colDir == 0) {
			for _, dir := range directions {
				if (dir.row != current.rowDir || dir.col != current.colDir) && (dir.row != -current.rowDir || dir.col != -current.colDir) {
					nextRow := current.row + dir.row
					nextCol := current.col + dir.col
					if 0 <= nextRow && nextRow < len(grid) && 0 <= nextCol && nextCol < len(grid[0]) {
						next := point{heat: current.heat + grid[nextRow][nextCol], row: nextRow, col: nextCol, rowDir: dir.row, colDir: dir.col, count: 1}
						queue = addToQueue(queue, next)
					}
				}
			}
		}
	}
}

func addToQueue(queue []point, value point) []point {
	var i int
	for i = 0; i < len(queue); i++ {
		if queue[i].heat > value.heat {
			break
		}
	}

	queue = append(queue[:i], append([]point{value}, queue[i:]...)...)

	return queue
}

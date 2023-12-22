package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Point struct {
	x     int
	y     int
	steps int
}

type Coord struct {
	x int
	y int
}

var directions = []Point{
	{0, 1, 0},
	{0, -1, 0},
	{1, 0, 0},
	{-1, 0, 0},
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	lines := strings.Split(string(input), "\n")

	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}

	start := Point{0, 0, 0}
	for x, row := range grid {
		for y, col := range row {
			if col == "S" {
				start = Point{x, y, 64}
			}
		}
	}

	seen := make(map[Coord]bool)
	seen[Coord{start.x, start.y}] = true
	queue := []Point{start}
	result := []Point{}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.steps%2 == 0 {
			result = append(result, item)
		}

		if item.steps == 0 {
			continue
		}

		for _, dir := range directions {
			next := Point{item.x + dir.x, item.y + dir.y, item.steps - 1}
			if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
				continue
			}
			if grid[next.x][next.y] == "#" {
				continue
			}
			if seen[Coord{next.x, next.y}] {
				continue
			}
			seen[Coord{next.x, next.y}] = true
			queue = append(queue, next)
		}
	}
	fmt.Println(len(result))
}

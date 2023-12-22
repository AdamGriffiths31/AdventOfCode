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
				start = Point{x, y, 26501365}
			}
		}
	}

	fill(grid, start)

	gridWidth := start.steps/len(grid) - 1
	odd := (gridWidth/2*2 + 1) * (gridWidth/2*2 + 1)
	even := ((gridWidth + 1) / 2 * 2) * ((gridWidth + 1) / 2 * 2)

	oddPoints := fill(grid, Point{start.x, start.y, len(grid)*2 + 1})
	evenPoints := fill(grid, Point{start.x, start.y, len(grid) * 2})
	fmt.Println(oddPoints, evenPoints)

	topCorner := fill(grid, Point{len(grid) - 1, start.y, len(grid) - 1})
	rightCorner := fill(grid, Point{start.x, 0, len(grid) - 1})
	bottomCorner := fill(grid, Point{0, start.y, len(grid) - 1})
	leftCorner := fill(grid, Point{start.x, len(grid) - 1, len(grid) - 1})
	fmt.Println(topCorner, rightCorner, bottomCorner, leftCorner)

	smallTopRight := fill(grid, Point{len(grid) - 1, 0, len(grid)/2 - 1})
	smallBottomRight := fill(grid, Point{0, 0, len(grid)/2 - 1})
	smallTopLeft := fill(grid, Point{len(grid) - 1, len(grid) - 1, len(grid)/2 - 1})
	smallBottomLeft := fill(grid, Point{0, len(grid) - 1, len(grid)/2 - 1})
	fmt.Println(smallTopRight, smallBottomRight, smallTopLeft, smallBottomLeft)

	largeTopRight := fill(grid, Point{len(grid) - 1, 0, len(grid)*3/2 - 1})
	largeBottomRight := fill(grid, Point{0, 0, len(grid)*3/2 - 1})
	largeTopLeft := fill(grid, Point{len(grid) - 1, len(grid) - 1, len(grid)*3/2 - 1})
	largeBottomLeft := fill(grid, Point{0, len(grid) - 1, len(grid)*3/2 - 1})
	fmt.Println(largeTopRight, largeBottomRight, largeTopLeft, largeBottomLeft)

	result := odd*oddPoints +
		even*evenPoints +
		topCorner + rightCorner + bottomCorner + leftCorner +
		((gridWidth + 1) * (smallTopRight + smallBottomRight + smallTopLeft + smallBottomLeft)) +
		(gridWidth * (largeTopRight + largeBottomRight + largeTopLeft + largeBottomLeft))
	fmt.Println(result)
}

func fill(grid [][]string, start Point) int {
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
	return len(result)
}

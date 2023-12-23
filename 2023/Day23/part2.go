package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Point struct {
	x int
	y int
}

type PointData struct {
	x int
	y int
	n int
}

var directions = []Point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

var seenGraph = map[Point]bool{}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	grid := []string{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		grid = append(grid, line)
	}

	idxStart := strings.Index(grid[0], ".")
	idxEnd := strings.Index(grid[len(grid)-1], ".")

	startP := Point{0, idxStart}
	endP := Point{len(grid) - 1, idxEnd}

	fmt.Printf("Start: %v End: %v\n", startP, endP)

	poi := []Point{startP, endP}
	for r, row := range grid {
		for c, col := range row {
			if col == '#' {
				continue
			}
			neighbors := 0
			for _, dir := range directions {
				neighbor := Point{r + dir.x, c + dir.y}
				if neighbor.x < 0 || neighbor.x >= len(grid) || neighbor.y < 0 || neighbor.y >= len(row) {
					continue
				}
				if grid[neighbor.x][neighbor.y] == '#' {
					continue
				}
				neighbors++
			}
			if neighbors >= 3 {
				poi = append(poi, Point{r, c})
			}
		}
	}
	fmt.Printf("POI: %v\n", poi)
	graph := map[Point]map[Point]int{}

	for _, p := range poi {
		stack := []PointData{{p.x, p.y, 0}}
		seen := map[Point]bool{}
		seen[p] = true
		graph[p] = map[Point]int{}

		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if current.n != 0 && contains(poi, Point{current.x, current.y}) {
				graph[p][Point{current.x, current.y}] = current.n
				continue
			}
			for _, dir := range directions {
				neighbor := Point{current.x + dir.x, current.y + dir.y}
				if neighbor.x < 0 || neighbor.x >= len(grid) || neighbor.y < 0 || neighbor.y >= len(grid[0]) {
					continue
				}
				if grid[neighbor.x][neighbor.y] == '#' {
					continue
				}
				if seen[neighbor] {
					continue
				}
				stack = append(stack, PointData{neighbor.x, neighbor.y, current.n + 1})
				seen[neighbor] = true
			}
		}
	}
	for k, v := range graph {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Printf("DFS: %v\n", dfs(startP, endP, graph))
}

func dfs(start, end Point, graph map[Point]map[Point]int) int {
	if start == end {
		return 0
	}

	m := math.MinInt32
	seenGraph[start] = true
	for next := range graph[start] {
		if seenGraph[next] {
			continue
		}
		m = max(m, dfs(next, end, graph)+graph[start][next])
	}
	seenGraph[start] = false

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func contains(poi []Point, p Point) bool {
	for _, v := range poi {
		if v == p {
			return true
		}
	}
	return false
}

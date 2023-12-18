package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

var directions = map[string]point{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	lines := strings.Split(string(input), "\n")

	points := []point{{0, 0}}
	boundaries := 0
	for _, line := range lines {
		path := strings.Split(line, " ")
		colour := path[2]
		colour = colour[2 : len(colour)-1]
		idx, _ := strconv.Atoi(string(colour[len(colour)-1]))
		newPoint := directions[string("RDLU"[idx])]
		steps, _ := strconv.ParseInt(colour[:len(colour)-1], 16, 64)
		boundaries += int(steps)
		lastPoint := points[len(points)-1]

		points = append(points, point{lastPoint.x + newPoint.x*int(steps), lastPoint.y + newPoint.y*int(steps)})
	}
	A := 0
	for i := 0; i < len(points); i++ {
		tmp := i - 1
		if tmp < 0 {
			tmp = 0
		}
		A += points[i].x * (points[tmp].y - points[(i+1)%len(points)].y)
	}
	fmt.Println(A)
	A = int(math.Abs(float64(A))) / 2
	fmt.Println(A)
	i := A - boundaries/2 + 1
	fmt.Println(i, A, boundaries)
	fmt.Println(i + boundaries)

}

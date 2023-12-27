package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type hailstone struct {
	startX    float64
	startY    float64
	startZ    float64
	velocityX float64
	velocityY float64
	velocityZ float64

	a float64
	b float64
	c float64
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	lines := strings.Split(string(input), "\n")

	hailstones := []hailstone{}
	for _, line := range lines {
		line = strings.Replace(line, "@", ",", -1)
		parts := strings.Split(line, ",")
		startX, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		startY, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		startZ, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		velocityX, _ := strconv.Atoi(strings.TrimSpace(parts[3]))
		velocityY, _ := strconv.Atoi(strings.TrimSpace(parts[4]))
		velocityZ, _ := strconv.Atoi(strings.TrimSpace(parts[5]))
		c := velocityY*startX - velocityX*startY
		h := hailstone{
			startX:    float64(startX),
			startY:    float64(startY),
			startZ:    float64(startZ),
			velocityX: float64(velocityX),
			velocityY: float64(velocityY),
			velocityZ: float64(velocityZ),
			a:         float64(velocityY),
			b:         float64(-velocityX),
			c:         float64(c),
		}
		hailstones = append(hailstones, h)
	}

	total := 0
	for i, hs1 := range hailstones {
		for _, hs2 := range hailstones[:i] {
			if hs1.a*hs2.b == hs1.b*hs2.a {
				continue
			}

			var x, y float64
			x = (hs1.c*hs2.b - hs2.c*hs1.b) / (hs1.a*hs2.b - hs2.a*hs1.b)
			y = (hs2.c*hs1.a - hs1.c*hs2.a) / (hs1.a*hs2.b - hs2.a*hs1.b)

			if (200000000000000 <= x && x <= 400000000000000) && (200000000000000 <= y && y <= 400000000000000) {
				if valid(hs1, x, y) && valid(hs2, x, y) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func valid(h hailstone, x, y float64) bool {
	if (x-h.startX)*h.velocityX >= 0 && (y-h.startY)*h.velocityY >= 0 {
		return true
	}
	return false

}

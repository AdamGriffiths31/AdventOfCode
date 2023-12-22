package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	bricks := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = strings.Replace(line, "~", ",", -1)
		parts := strings.Split(line, ",")
		arr := []int{}
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			arr = append(arr, num)
		}
		bricks = append(bricks, arr)
	}
	bricks = sortBricks(bricks)

	for i, brick := range bricks {
		maxZ := 1
		for _, checkBrick := range bricks[:i] {
			if overlaps(brick, checkBrick) {
				maxZ = max(maxZ, checkBrick[5]+1)
			}
		}
		bricks[i][5] -= bricks[i][2] - maxZ
		bricks[i][2] = maxZ
	}

	bricks = sortBricks(bricks)

	supportsV := make(map[int][]int, len(bricks))
	supportsK := make(map[int][]int, len(bricks))

	for j, upperBrick := range bricks {
		for i, lowerBrick := range bricks[:j] {
			if overlaps(lowerBrick, upperBrick) && upperBrick[2] == lowerBrick[5]+1 {
				supportsV[i] = append(supportsV[i], j)
				supportsK[j] = append(supportsK[j], i)
			}
		}
	}

	total := 0
	for i := range bricks {
		count := 0
		for _, support := range supportsV[i] {
			if len(supportsK[support]) >= 2 {
				count++
			}
		}
		if count == len(supportsV[i]) {
			total++
		}
	}

	fmt.Println(total)
}

func sortBricks(bricks [][]int) [][]int {
	for i := 0; i < len(bricks); i++ {
		for j := i + 1; j < len(bricks); j++ {
			if bricks[j][2] < bricks[i][2] {
				bricks[i], bricks[j] = bricks[j], bricks[i]
			}
		}
	}
	return bricks
}

func overlaps(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[3], b[3]) && max(a[1], b[1]) <= min(a[4], b[4])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileData, _ := ioutil.ReadFile("input.txt")

	seedsField := strings.Split(string(fileData), "\n")[0]
	inputs := getSeeds(seedsField)

	seeds := [][]int{}
	for i := 0; i < len(inputs); i += 2 {
		seeds = append(seeds, []int{inputs[i], inputs[i] + inputs[i+1]})
	}

	blocks := strings.Split(string(fileData), "\n\n")[1:]
	for _, block := range blocks {
		ranges := make([][]int, 0)
		for _, line := range strings.Split(block, "\n")[1:] {
			if line == "" {
				continue
			}
			ranges = append(ranges, splitLine(line))
		}
		newSlice := []int{}
		for _, seed := range inputs {
			added := false
			for _, num := range ranges {
				if num[1] <= seed && seed < num[1]+num[2] {
					newSlice = append(newSlice, seed-num[1]+num[0])
					added = true
					break
				}
			}
			if !added {
				newSlice = append(newSlice, seed)
			}
		}
		inputs = newSlice

	}
	fmt.Printf("Seeds: %+v\n", inputs)
	fmt.Printf("Min: %+v\n", minList(inputs))
	part2()
}

func part2() {
	fileData, _ := ioutil.ReadFile("input.txt")

	seedsField := strings.Split(string(fileData), "\n")[0]
	inputs := getSeeds(seedsField)

	seeds := [][]int{}
	for i := 0; i < len(inputs); i += 2 {
		seeds = append(seeds, []int{inputs[i], inputs[i] + inputs[i+1]})
	}

	blocks := strings.Split(string(fileData), "\n\n")[1:]
	for _, block := range blocks {
		ranges := make([][]int, 0)
		for _, line := range strings.Split(block, "\n")[1:] {
			if line == "" {
				continue
			}
			ranges = append(ranges, splitLine(line))
		}
		newSlice := [][]int{}
		for len(seeds) > 0 {
			s, e := popSeed(&seeds)
			appended := false
			for _, num := range ranges {
				overlapStart := max(s, num[1])
				overlapEnd := min(e, num[1]+num[2])
				if overlapStart < overlapEnd {
					appended = true
					newSlice = append(newSlice, []int{overlapStart - num[1] + num[0], overlapEnd - num[1] + num[0]})
					if overlapStart > s {
						seeds = append(seeds, []int{s, overlapStart})
					}
					if e > overlapEnd {
						seeds = append(seeds, []int{overlapEnd, e})
					}
					break
				}
			}
			if !appended {
				newSlice = append(newSlice, []int{s, e})
			}
		}
		seeds = newSlice
	}
	fmt.Printf("Seeds: %+v\n", seeds)
	fmt.Printf("Min: %+v\n", minList2(seeds))
}

func popSeed(seeds *[][]int) (int, int) {
	if len(*seeds) == 0 {
		return 0, 0
	}

	lastIdx := len(*seeds) - 1
	seed := (*seeds)[lastIdx]
	*seeds = (*seeds)[:lastIdx]
	return seed[0], seed[1]
}

func splitLine(line string) []int {
	numbers := []int{}
	for _, str := range strings.Split(line, " ") {
		if str == "" {
			continue
		}

		numInt, _ := strconv.Atoi(string(str))
		numbers = append(numbers, numInt)
	}
	return numbers
}

func getSeeds(seedsField string) []int {
	seeds := []int{}
	seedsField = strings.Split(seedsField, ":")[1]
	for _, str := range strings.Split(seedsField, " ") {
		if str == "" {
			continue
		}

		seedInt, _ := strconv.Atoi(string(str))
		seeds = append(seeds, seedInt)
	}
	return seeds
}

func minList2(list [][]int) []int {
	min := list[0]
	for _, num := range list {
		if num[0] < min[0] {
			min = num
		}
	}
	return min
}

func minList(list []int) int {
	min := list[0]
	for _, num := range list {
		if num < min {
			min = num
		}
	}
	return min
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

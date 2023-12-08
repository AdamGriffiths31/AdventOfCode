package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strings"
)

type Node struct {
	Name  string
	Left  string
	Right string
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	parts := strings.Split(string(input), "\n")
	directions := parts[0]
	fmt.Printf("Directions: %v\n", directions)

	nodeList := parts[1:]
	nodes := make(map[string]Node)
	for _, node := range nodeList {
		if node == "" {
			continue
		}
		pattern := `(\w+)\s*=\s*\((\w+),\s*(\w+)\)`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(node)

		nodes[matches[1]] = Node{Name: matches[1], Left: matches[2], Right: matches[3]}
	}

	startingNodes := []Node{}
	for name, node := range nodes {
		if name[2] == 'A' {
			startingNodes = append(startingNodes, node)
		}
	}

	zList := []int{}
	for _, startingNode := range startingNodes {
		count := 0
		for startingNode.Name[2] != 'Z' {
			nextNode := getNextNode(startingNode.Name, nodes, getDirections(directions, count))
			startingNode = nodes[nextNode]
			count++
		}
		zList = append(zList, count)
	}
	fmt.Printf("LCM: %v\n", findLCMOfList(zList))
}

func getNextNode(current string, nodes map[string]Node, instruction string) string {
	currentNode := nodes[current]

	if instruction == "L" {
		return currentNode.Left
	}
	return currentNode.Right
}

func getDirections(directions string, count int) string {
	if count >= len(directions) {
		return string(directions[count%len(directions)])
	}
	return string(directions[count])
}

func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findLCM(a, b int) int {
	gcd := findGCD(a, b)
	lcm := int(math.Abs(float64(a*b))) / gcd
	return lcm
}

func findLCMOfList(nums []int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = findLCM(lcm, nums[i])
	}
	return lcm
}

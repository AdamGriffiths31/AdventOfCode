package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	parts := strings.Split(string(input), "\n")
	directions := parts[0]

	nodeList := parts[1:]
	nodes := make(map[string]Node)
	for _, node := range nodeList {
		if node == "" {
			continue
		}
		pattern := `(\w+)\s*=\s*\((\w+),\s*(\w+)\)`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(node)

		nodes[matches[1]] = Node{Left: matches[2], Right: matches[3]}
	}

	current := "AAA"
	destination := "ZZZ"
	count := 0
	for current != destination {
		direction := getDirections(directions, count)
		current = getNextNode(current, nodes, direction)
		count++
	}
	fmt.Printf("count: %d\n", count)
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

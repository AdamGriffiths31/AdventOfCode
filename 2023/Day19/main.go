package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type comparisonFunc func(a, b int) bool

var ops = map[string]comparisonFunc{
	">": func(a, b int) bool { return a > b },
	"<": func(a, b int) bool { return a < b },
}

type data struct {
	rules    []dataRule
	fallback string
}

type dataRule struct {
	key    string
	cmp    string
	n      int
	target string
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	blocks := strings.Split(string(input), "\n\n")
	lines := strings.Split(blocks[0], "\n")

	workflow := make(map[string]data)
	for _, line := range lines {
		name := strings.Split(line, "{")[0]
		rest := strings.Trim(strings.Split(line, "{")[1], "}")
		fallbackList := strings.Split(rest, ",")
		fallback := fallbackList[len(fallbackList)-1]
		rulesStr := strings.TrimSuffix(rest, fallback)
		rules := strings.Split(rulesStr, ",")

		rulesList := []dataRule{}
		for _, rule := range rules {
			if rule == "" {
				continue
			}
			parts := strings.Split(rule, ":")
			comparison := parts[0]
			target := parts[1]
			key := comparison[0]
			cmp := comparison[1]
			n, _ := strconv.Atoi(comparison[2:])
			datarule := dataRule{key: string(key), cmp: string(cmp), n: n, target: target}
			rulesList = append(rulesList, datarule)
		}
		data := data{rules: rulesList, fallback: fallback}
		workflow[name] = data
	}

	total := 0
	lines = strings.Split(blocks[1], "\n")
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		items := make(map[string]int)
		parts := strings.Split(strings.Trim(line[1:], "}"), ",")
		for _, part := range parts {
			partsInner := strings.Split(part, "=")
			char := partsInner[0]
			n, _ := strconv.Atoi(partsInner[1])
			items[char] = n
		}
		if acceptItem(items, "in", workflow) {
			tmp := 0
			for val := range items {
				tmp += items[val]
			}
			total += tmp
		}
	}
	println(total)
}

func acceptItem(item map[string]int, key string, workflows map[string]data) bool {
	if key == "R" {
		return false
	}
	if key == "A" {
		return true
	}
	data := workflows[key]

	for _, rule := range data.rules {
		if ops[rule.cmp](item[rule.key], rule.n) {
			return acceptItem(item, rule.target, workflows)
		}
	}
	return acceptItem(item, data.fallback, workflows)
}

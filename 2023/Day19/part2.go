package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

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

type rangeValue struct {
	min int
	max int
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

	ranges := make(map[string]rangeValue)
	keys := []string{"x", "m", "a", "s"}
	for _, key := range keys {
		ranges[key] = rangeValue{1, 4000}
	}

	result := count("in", ranges, workflow)
	println(result)
}

func count(name string, ranges map[string]rangeValue, workflows map[string]data) int {
	if name == "R" {
		return 0
	}
	if name == "A" {
		product := 1
		for _, r := range ranges {
			product *= r.max - r.min + 1
		}
		return product
	}

	total := 0
	data := workflows[name]
	for _, rule := range data.rules {
		rangeVal := ranges[rule.key]
		var T, F rangeValue
		if rule.cmp == "<" {
			T = rangeValue{min: rangeVal.min, max: min(rule.n-1, rangeVal.max)}
			F = rangeValue{min: max(rule.n, rangeVal.min), max: rangeVal.max}
		} else {
			T = rangeValue{min: max(rule.n+1, rangeVal.min), max: rangeVal.max}
			F = rangeValue{min: rangeVal.min, max: min(rule.n, rangeVal.max)}
		}

		if T.min <= T.max {
			ranges[rule.key] = T
			total += count(rule.target, ranges, workflows)
		}
		if F.min <= F.max {
			ranges[rule.key] = F
		} else {
			break
		}
	}

	if total == 0 {
		total += count(data.fallback, ranges, workflows)
	}

	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

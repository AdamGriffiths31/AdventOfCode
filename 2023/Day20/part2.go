package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type module struct {
	Name         string
	Type         string
	Outputs      []string
	Memory       interface{}
	MemoryString string
}

type queueObject struct {
	Origin string
	Target string
	Pulse  string
}

var targets = []string{}
var modules = map[string]module{}

func (m module) getMemory() interface{} {
	if m.Type == "%" {
		return "off"
	}
	return m.Memory
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	input = input[:len(input)-1]
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		name := parts[0]
		action := parts[1]

		if name == "broadcaster" {
			targets = append(targets, strings.Split(action, ", ")...)
		} else {
			mod := module{
				Name:         name[1:],
				Outputs:      strings.Split(action, ", "),
				Type:         string(name[0]),
				Memory:       make(map[string]string),
				MemoryString: "off",
			}
			modules[name[1:]] = mod
		}
	}

	for _, mod := range modules {
		for _, output := range mod.Outputs {
			if contains(output, modules) && modules[output].Type == "&" {
				modules[output].Memory.(map[string]string)[mod.Name] = "lo"
			}
		}
	}
	queue := []queueObject{}
	feed := module{}
	seen := map[string]int{}
	for _, mod := range modules {
		for _, output := range mod.Outputs {
			if output == "rx" {
				feed = mod
			}
		}
	}

	cycleLengths := map[string]string{}
	for _, mod := range modules {
		for _, output := range mod.Outputs {
			if output == feed.Name {
				cycleLengths[mod.Name] = "None"
			}
		}
	}
	fmt.Println("cycleLengths: ", cycleLengths)
	for _, mod := range modules {
		for _, output := range mod.Outputs {
			if output == feed.Name {
				seen[mod.Name] = 0
			}
		}
	}

	buttonPresses := 0
	for true {
		buttonPresses++
		for _, target := range targets {
			queue = append(queue, queueObject{"broadcaster", target, "lo"})
		}

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]

			if !contains(item.Target, modules) {
				continue
			}

			mod := modules[item.Target]
			if mod.Name == feed.Name && item.Pulse == "hi" {
				seen[item.Origin]++
				if !containsNonNil(item.Origin, cycleLengths) {
					fmt.Println("seen: ", seen)
					cycleLengths[item.Origin] = fmt.Sprintf("%d", buttonPresses)
				}
				if allSeen(seen) {
					fmt.Println("cycleLengths: ", cycleLengths)
					ints := []int{}
					for _, value := range cycleLengths {
						num, _ := strconv.Atoi(value)
						ints = append(ints, num)
					}
					lcm := findLCMOfList(ints)
					fmt.Println("lcm: ", lcm)
					return
				}
			}

			if mod.Type == "%" {
				if item.Pulse == "lo" {
					if mod.MemoryString == "off" {
						mod.MemoryString = "on"
					} else {
						mod.MemoryString = "off"
					}
					outgoing := "lo"
					if mod.MemoryString == "on" {
						outgoing = "hi"
					}
					for _, output := range mod.Outputs {
						queue = append(queue, queueObject{mod.Name, output, outgoing})
					}
				}
			} else {
				mod.Memory.(map[string]string)[item.Origin] = item.Pulse
				isHigh := areAllHigh(mod.Memory.(map[string]string))
				outgoing := "hi"
				if isHigh {
					outgoing = "lo"
				}
				for _, output := range mod.Outputs {
					queue = append(queue, queueObject{mod.Name, output, outgoing})
				}
			}
			modules[item.Target] = mod
		}
	}
}

func allSeen(seen map[string]int) bool {
	fmt.Println("seen: ", seen)
	for _, value := range seen {
		if value == 0 {
			return false
		}
	}
	return true
}

func areAllHigh(outputs map[string]string) bool {
	for _, output := range outputs {
		if output == "lo" {
			return false
		}
	}
	return true
}

func containsNonNil(item string, list map[string]string) bool {
	item, found := list[item]
	if found && item != "None" {
		return true
	}
	return false
}

func contains(item string, list map[string]module) bool {
	_, found := list[item]
	return found
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

package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Println(modules)

	lowPulse := 0
	highPulse := 0
	queue := []queueObject{}

	for i := 0; i < 1000; i++ {
		lowPulse++
		for _, target := range targets {
			queue = append(queue, queueObject{"broadcaster", target, "lo"})
		}

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]

			if item.Pulse == "hi" {
				highPulse++
			} else {
				lowPulse++
			}

			if !contains(item.Target, modules) {
				continue
			}

			mod := modules[item.Target]
			if mod.Type == "%" {
				if item.Pulse == "lo" {
					fmt.Println("memory: ", mod.MemoryString)
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

	fmt.Println("Lowest pulse:", lowPulse)
	fmt.Println("Highest pulse:", highPulse)
	fmt.Println(lowPulse * highPulse)
}

func areAllHigh(outputs map[string]string) bool {
	for _, output := range outputs {
		if output == "lo" {
			return false
		}
	}
	return true
}

func contains(item string, list map[string]module) bool {
	_, found := list[item]
	return found
}

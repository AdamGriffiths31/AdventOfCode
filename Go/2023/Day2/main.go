package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var colors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	result := 0
	result2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += partOne(scanner.Text())
		result2 += partTwo(scanner.Text())
	}
	fmt.Println(result)
	fmt.Println(result2)
}

func partOne(input string) int {
	game := strings.Split(input, ":")[0]
	gameNuber := strings.Split(game, " ")[1]
	gameParts := strings.Split(input, ":")[1]
	games := strings.Split(gameParts, ";")
	for _, game := range games {
		red := 0
		green := 0
		blue := 0
		individualGame := strings.Split(game, ",")
		for _, cube := range individualGame {
			cube = strings.TrimSpace(cube)
			number, _ := strconv.Atoi(strings.Split(cube, " ")[0])
			color := strings.Split(cube, " ")[1]
			switch color {
			case "red":
				red += number
			case "green":
				green += number
			case "blue":
				blue += number
			default:
				panic("unknown color")
			}
			if red > colors["red"] || green > colors["green"] || blue > colors["blue"] {
				return 0
			}
		}
	}
	returnVal, _ := strconv.Atoi(gameNuber)
	return int(returnVal)
}

func partTwo(input string) int {
	gameparts := strings.Split(input, ":")[1]
	games := strings.Split(gameparts, ";")
	red := 0
	green := 0
	blue := 0
	for _, game := range games {
		individualGame := strings.Split(game, ",")
		for _, cube := range individualGame {
			cube = strings.TrimSpace(cube)
			number, _ := strconv.Atoi(strings.Split(cube, " ")[0])
			color := strings.Split(cube, " ")[1]
			switch color {
			case "red":
				if number > red {
					red = number
				}
			case "green":
				if number > green {
					green = number
				}
			case "blue":
				if number > blue {
					blue = number
				}
			}
		}
	}
	fmt.Printf("red: %d, green: %d, blue: %d\n", red, green, blue)
	return red * green * blue
}

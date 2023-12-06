package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileData, _ := ioutil.ReadFile("input.txt")
	time := strings.Split(string(fileData), "\n")[0]
	distances := strings.Split(string(fileData), "\n")[1]

	timeList := convertToList(time)
	distancesList := convertToList(distances)

	result := 1
	for i, time := range timeList {
		waysToWin := 0
		for timePressed := 1; timePressed < time; timePressed++ {
			timeRacing := time - timePressed
			distance := timeRacing * timePressed
			if distance > distancesList[i] {
				waysToWin++
			}
		}
		result *= waysToWin
	}
	fmt.Printf("Result: %v\n", result)
	part2(string(fileData))
}

func part2(input string) {
	time := strings.Split(input, "\n")[0]
	distances := strings.Split(input, "\n")[1]

	timeInt := convertToSingleInt(time)
	distancesInt := convertToSingleInt(distances)

	waysToWin := 0
	for timePressed := 1; timePressed < timeInt; timePressed++ {
		timeRacing := timeInt - timePressed
		distance := timeRacing * timePressed
		if distance > distancesInt {
			waysToWin++
		}
	}
	fmt.Printf("Result: %v\n", waysToWin)

}

func convertToSingleInt(input string) int {
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(input, -1)

	var result string
	for _, match := range matches {
		result += match
	}

	resultInt, _ := strconv.Atoi(result)
	return resultInt
}

func convertToList(input string) []int {
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(input, -1)

	var result []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			result = append(result, num)
		}
	}

	return result
}

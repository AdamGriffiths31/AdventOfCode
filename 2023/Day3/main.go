package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	engine := make(map[int]string)
	row := 0
	for scanner.Scan() {
		engine[row] = scanner.Text()
		row++
	}
	part1(engine)
}

func part1(engine map[int]string) {
	score := 0
	gearRatio := 0
	for i := 0; i < len(engine); i++ {
		for j := 0; j < len(engine[i]); j++ {
			if !isDigit(rune(engine[i][j])) && engine[i][j] != '.' {
				//check diganolly up and left
				scoreToAdd := []int{}
				if i > 0 && j > 0 && isDigit(rune(engine[i-1][j-1])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i-1, j-1))
				}
				//check diganolly up and right
				if i > 0 && j < len(engine[i])-1 && isDigit(rune(engine[i-1][j+1])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i-1, j+1))
				}
				//check diganolly down and left
				if i < len(engine)-1 && j > 0 && isDigit(rune(engine[i+1][j-1])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i+1, j-1))
				}
				//check diganolly down and right
				if i < len(engine)-1 && j < len(engine[i])-1 && isDigit(rune(engine[i+1][j+1])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i+1, j+1))
				}
				//check up
				if i > 0 && isDigit(rune(engine[i-1][j])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i-1, j))
				}
				//check down
				if i < len(engine)-1 && isDigit(rune(engine[i+1][j])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i+1, j))
				}
				//check left
				if j > 0 && isDigit(rune(engine[i][j-1])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i, j-1))
				}
				//check right
				if j < len(engine[i])-1 && isDigit(rune(engine[i][j+1])) {
					scoreToAdd = append(scoreToAdd, findFullNumber(engine, i, j+1))
				}
				if engine[i][j] == '*' && len(scoreToAdd) == 2 {
					gearRatio += scoreToAdd[0] * scoreToAdd[1]
				}
				for _, val := range scoreToAdd {
					score += val
				}
			}
		}
	}
	fmt.Printf("score: %d\n", score)
	fmt.Printf("gear ratio: %d\n", gearRatio)
}

func findFullNumber(engine map[int]string, i, j int) int {
	var number string
	start := j
	for start > 0 && isDigit(rune(engine[i][start-1])) {
		start--
	}
	for start < len(engine[i]) && isDigit(rune(engine[i][start])) {
		number += string(engine[i][start])
		//remove the number from the engine
		engine[i] = engine[i][:start] + "." + engine[i][start+1:]
		start++
	}
	fmt.Printf("found number: %s\n", number)
	returnVal, _ := strconv.Atoi(number)
	return returnVal
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

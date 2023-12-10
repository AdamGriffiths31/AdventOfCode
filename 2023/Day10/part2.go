package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid []string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	var sr, sc int
	seen := make(map[[2]int]bool)
	var q = make([][2]int, 0)

	for r, row := range grid {
		for c, ch := range row {
			if ch == 'S' {
				sr, sc = r, c
				break
			}
		}
	}

	seen[[2]int{sr, sc}] = true
	q = append(q, [2]int{sr, sc})

	maybeS := map[string]struct{}{"|": {}, "-": {}, "J": {}, "L": {}, "7": {}, "F": {}}

	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		ch := grid[r][c]

		if r > 0 && (ch == 'S' || ch == '|' || ch == 'J' || ch == 'L') && (grid[r-1][c] == '|' || grid[r-1][c] == '7' || grid[r-1][c] == 'F') && !seen[[2]int{r - 1, c}] {
			seen[[2]int{r - 1, c}] = true
			q = append(q, [2]int{r - 1, c})
			if ch == 'S' {
				intersectionSet := map[string]struct{}{"|": {}, "J": {}, "L": {}}
				maybeS = removeIntersection(maybeS, intersectionSet)
			}
		}

		if r < len(grid)-1 && (ch == 'S' || ch == '|' || ch == '7' || ch == 'F') && (grid[r+1][c] == '|' || grid[r+1][c] == 'J' || grid[r+1][c] == 'L') && !seen[[2]int{r + 1, c}] {
			seen[[2]int{r + 1, c}] = true
			q = append(q, [2]int{r + 1, c})
			if ch == 'S' {
				intersectionSet := map[string]struct{}{"|": {}, "7": {}, "F": {}}
				maybeS = removeIntersection(maybeS, intersectionSet)
			}
		}

		if c > 0 && (ch == 'S' || ch == '-' || ch == 'J' || ch == '7') && (grid[r][c-1] == '-' || grid[r][c-1] == 'L' || grid[r][c-1] == 'F') && !seen[[2]int{r, c - 1}] {
			seen[[2]int{r, c - 1}] = true
			q = append(q, [2]int{r, c - 1})
			if ch == 'S' {
				intersectionSet := map[string]struct{}{"-": {}, "J": {}, "7": {}}
				maybeS = removeIntersection(maybeS, intersectionSet)
			}
		}

		if c < len(grid[r])-1 && (ch == 'S' || ch == '-' || ch == 'L' || ch == 'F') && (grid[r][c+1] == '-' || grid[r][c+1] == 'J' || grid[r][c+1] == '7') && !seen[[2]int{r, c + 1}] {
			seen[[2]int{r, c + 1}] = true
			q = append(q, [2]int{r, c + 1})
			if ch == 'S' {
				intersectionSet := map[string]struct{}{"-": {}, "L": {}, "F": {}}
				maybeS = removeIntersection(maybeS, intersectionSet)
			}
		}
	}

	grid[sr] = grid[sr][:sc] + getSvalues(maybeS) + grid[sr][sc+1:]
	for r, row := range grid {
		for c := range row {
			if seen[[2]int{r, c}] {
				fmt.Printf("%v", string(grid[r][c]))
			} else {
				grid[r] = grid[r][:c] + "." + grid[r][c+1:]
				fmt.Printf("%v", string(grid[r][c]))
			}
		}
		fmt.Println()
	}

	outside := make(map[[2]int]bool)
	for r, row := range grid {
		var within bool
		var up bool

		for c, ch := range row {
			fmt.Printf("%v\n", string(ch))
			fmt.Printf("up: %v\n", up)
			switch ch {
			case '|':
				within = !within
			case '-':
			case 'L', 'F':
				up = ch == 'L'
				fmt.Printf("change up: %v\n", up)
			case '7', 'J':
				temp := conditonalValue(up)
				if ch != temp {
					within = !within
				}
				up = false
			case '.':
			default:
				log.Panicf("Unexpected character (horizontal): %c", ch)
			}

			if !within {
				outside[[2]int{r, c}] = true
			}
		}
	}

	//print grid
	fmt.Println()
	for r, row := range grid {
		for c := range row {
			if outside[[2]int{r, c}] {
				fmt.Printf("%v", "#")
			} else {
				fmt.Printf("%v", ".")
			}
		}
		fmt.Println()
	}
	result := len(grid)*len(grid[0]) - len(union(outside, seen))

	// Print the result
	fmt.Println(result)
}

func conditonalValue(up bool) rune {
	if up {
		return 'J'
	}
	return '7'
}

func getSvalues(values map[string]struct{}) string {
	var s string
	for key := range values {
		s += key
	}
	return s
}
func removeIntersection(values, intersect map[string]struct{}) map[string]struct{} {
	for key := range values {
		if _, exists := intersect[key]; !exists {
			fmt.Printf("removing %v\n", key)
			delete(values, key)
		}
	}
	return values
}

func union(a, b map[[2]int]bool) map[[2]int]struct{} {
	result := make(map[[2]int]struct{})

	for k := range a {
		result[k] = struct{}{}
	}

	for k := range b {
		result[k] = struct{}{}
	}

	return result
}

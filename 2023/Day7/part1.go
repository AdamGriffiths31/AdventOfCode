package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var letters = map[string]string{
	"T": "A",
	"J": "B",
	"Q": "C",
	"K": "D",
	"A": "E",
}

type Hand struct {
	cards          string
	originalCards  string
	bid            int
	classification int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	hands := []Hand{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		cards := parts[0]
		bid := parts[1]
		bidInt, _ := strconv.Atoi(bid)
		hands = append(hands, Hand{cards, cards, bidInt, 0})
	}
	strength(hands)
}

func strength(hand []Hand) {
	for i, h := range hand {
		var replacedCards strings.Builder
		for _, c := range h.cards {
			if replacement, ok := letters[string(c)]; ok {
				replacedCards.WriteString(replacement)
			} else {
				replacedCards.WriteRune(c)
			}
		}
		hand[i].cards = replacedCards.String()
	}
	for i, h := range hand {
		hand[i].classification = classify(h)
	}

	sort.SliceStable(hand, func(i, j int) bool {
		if hand[i].classification != hand[j].classification {
			return hand[i].classification < hand[j].classification
		}
		return hand[i].cards < hand[j].cards
	})

	total := 0
	for i, h := range hand {
		fmt.Printf("%s %d\n", h.originalCards, h.classification)
		total += (i + 1) * h.bid
	}
	fmt.Println(total)
}

func classify(hand Hand) int {
	counts := count(hand.cards)

	if contains(counts, 5) {
		return 6
	}
	if contains(counts, 4) {
		return 5
	}
	if contains(counts, 3) && contains(counts, 2) {
		return 4
	}
	if contains(counts, 3) {
		return 3
	}
	if countOccurrences(counts, 2) == 4 {
		return 2
	}
	if contains(counts, 2) {
		return 1
	}

	return 0
}

func count(card string) []int {
	counts := map[string]int{}
	for _, c := range card {
		counts[string(c)]++
	}

	returnVal := []int{}
	for _, c := range card {
		returnVal = append(returnVal, counts[string(c)])
	}
	return returnVal
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func countOccurrences(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}

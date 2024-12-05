package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("4/input.txt")

	if err != nil {
		panic(err)
	}

	wordSearch := make([][]rune, 0)
	row := make([]rune, 0)
	for _, c := range string(input) {
		if c == '\n' {
			wordSearch = append(wordSearch, row)
			row = make([]rune, 0)
		} else {
			row = append(row, c)
		}
	}
	wordSearch = append(wordSearch, row)

	part1Search(wordSearch)

	part2Search(wordSearch)
}

func part2Search(wordSearch [][]rune) {
	tally := 0

	for y, yChars := range wordSearch {
		for x, xChar := range yChars {
			if xChar == 'A' {
				//Check constraints
				if x-1 < 0 {
					continue
				} else if x+1 >= len(wordSearch[y]) {
					continue
				} else if y-1 < 0 {
					continue
				} else if y+1 >= len(wordSearch[y]) {
					continue
				}

				posTally := 0
				posTally += part2SearchLeft(wordSearch, x, y)
				posTally += part2SearchRight(wordSearch, x, y)

				if posTally != 2 {
					continue
				}

				fmt.Printf("%d, %d: %d\n", x, y, posTally)
				tally++
			}
		}
	}

	fmt.Printf("Total: %d\n", tally)
}

func part2SearchLeft(wordSearch [][]rune, x int, y int) int {
	tally := 0

	if wordSearch[y-1][x-1] == 'M' {
		if wordSearch[y+1][x+1] == 'S' {
			tally++
		}
	}

	if wordSearch[y+1][x-1] == 'M' {
		if wordSearch[y-1][x+1] == 'S' {
			tally++
		}
	}

	return tally
}

func part2SearchRight(wordSearch [][]rune, x int, y int) int {
	tally := 0

	if wordSearch[y-1][x+1] == 'M' {
		if wordSearch[y+1][x-1] == 'S' {
			tally++
		}
	}

	if wordSearch[y+1][x+1] == 'M' {
		if wordSearch[y-1][x-1] == 'S' {
			tally++
		}
	}

	return tally
}

package main

import (
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

	part1Search(wordSearch, "XMAS")

	part2Search(wordSearch)
}

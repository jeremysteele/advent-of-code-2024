package main

import "fmt"

var searchDirections = [][]int{
	{0, -1},  //Up
	{0, 1},   //Down
	{-1, 0},  //Left
	{1, 0},   //Right
	{-1, -1}, // Up Left
	{1, -1},  //Up Right
	{-1, 1},  // Down left
	{1, 1},   // Down right
}

func part1Search(wordSearch [][]rune, word string) {
	tally := 0

	for y, yChars := range wordSearch {
		for x, xChar := range yChars {
			if xChar == rune(word[0]) {
				tally += searchFromPos(wordSearch, word, x, y)
			}
		}
	}

	fmt.Printf("Total: %d\n", tally)
}

func searchFromPos(wordSearch [][]rune, word string, x int, y int) int {
	tally := 0

	boardWidth := len(wordSearch[y])
	boardHeight := len(wordSearch)

	for _, d := range searchDirections {
		match := true

		for i, c := range word {
			searchX := x + d[0]*i
			searchY := y + d[1]*i

			// Constraint check
			if searchX < 0 || searchX >= boardWidth {
				match = false
				break
			} else if searchY < 0 || searchY >= boardHeight {
				match = false
				break
			}

			if c != wordSearch[searchY][searchX] {
				match = false
				break
			}
		}

		if match {
			tally++
		}
	}

	fmt.Printf("%d, %d: %d\n", x, y, tally)

	return tally
}

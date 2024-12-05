package main

import "fmt"

func part1Search(wordSearch [][]rune) {
	tally := 0

	for y, yChars := range wordSearch {
		for x, xChar := range yChars {
			if xChar == 'X' {
				tally += searchFromPos(wordSearch, x, y)
			}
		}
	}

	fmt.Printf("Total: %d\n", tally)
}

func searchLeft(wordSearch [][]rune, x int, y int) int {
	tally := 0

	// Can we even search left?
	if x-3 < 0 {
		return 0
	}

	if wordSearch[y][x-1] == 'M' {
		if wordSearch[y][x-2] == 'A' {
			if wordSearch[y][x-3] == 'S' {
				tally++
			}
		}
	}

	if y-3 >= 0 {
		if wordSearch[y-1][x-1] == 'M' {
			if wordSearch[y-2][x-2] == 'A' {
				if wordSearch[y-3][x-3] == 'S' {
					tally++
				}
			}
		}
	}

	if y+3 < len(wordSearch) {
		if wordSearch[y+1][x-1] == 'M' {
			if wordSearch[y+2][x-2] == 'A' {
				if wordSearch[y+3][x-3] == 'S' {
					tally++
				}
			}
		}
	}

	return tally
}

func searchRight(wordSearch [][]rune, x int, y int) int {
	tally := 0

	// Can we even search left?
	if x+3 >= len(wordSearch[y]) {
		return 0
	}

	if wordSearch[y][x+1] == 'M' {
		if wordSearch[y][x+2] == 'A' {
			if wordSearch[y][x+3] == 'S' {
				tally++
			}
		}
	}

	if y-3 >= 0 {
		if wordSearch[y-1][x+1] == 'M' {
			if wordSearch[y-2][x+2] == 'A' {
				if wordSearch[y-3][x+3] == 'S' {
					tally++
				}
			}
		}
	}

	if y+3 < len(wordSearch) {
		if wordSearch[y+1][x+1] == 'M' {
			if wordSearch[y+2][x+2] == 'A' {
				if wordSearch[y+3][x+3] == 'S' {
					tally++
				}
			}
		}
	}

	return tally
}

func searchUp(wordSearch [][]rune, x int, y int) int {
	// Can we search up?
	if y-3 < 0 {
		return 0
	}

	if wordSearch[y-1][x] == 'M' {
		if wordSearch[y-2][x] == 'A' {
			if wordSearch[y-3][x] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func searchDown(wordSearch [][]rune, x int, y int) int {
	// Can we search down?
	if y+3 >= len(wordSearch) {
		return 0
	}

	if wordSearch[y+1][x] == 'M' {
		if wordSearch[y+2][x] == 'A' {
			if wordSearch[y+3][x] == 'S' {
				return 1
			}
		}
	}

	return 0
}

func searchFromPos(wordSearch [][]rune, x int, y int) int {
	tally := 0

	tally += searchLeft(wordSearch, x, y)
	tally += searchRight(wordSearch, x, y)
	tally += searchUp(wordSearch, x, y)
	tally += searchDown(wordSearch, x, y)

	fmt.Printf("%d, %d: %d\n", x, y, tally)

	return tally
}

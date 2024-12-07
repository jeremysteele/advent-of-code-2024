package main

import (
	"fmt"
	"os"
	"strings"
)

// They want game? Let's make a friggin game!

type vector struct {
	x, y int
}

type position vector

var directions = []vector{
	vector{0, -1}, // Up
	vector{1, 0},  // Right
	vector{0, 1},  // Down
	vector{-1, 0}, // Left
}

var maxDirection = len(directions)

type puzzleInstance struct {
	guardPosition position

	guardVector int

	board [][]rune

	boardHeight int
	boardWidth  int
}

func (puzzle *puzzleInstance) isPositionInbounds(pos position) bool {
	if pos.x >= 0 && pos.x < puzzle.boardWidth {
		if pos.y >= 0 && pos.y < puzzle.boardHeight {
			return true
		}
	}

	return false
}

func (puzzle *puzzleInstance) tabulateMoves() int {
	total := 0

	for _, y := range puzzle.board {
		for _, r := range y {
			if r == 'X' {
				total++
			}
		}
	}

	return total
}

func (puzzle *puzzleInstance) render() {
	renderHeight := 20
	renderWidth := 20

	fmt.Printf("-----------------------------------\n")
	for y := puzzle.guardPosition.y - (renderHeight / 2); y < puzzle.guardPosition.y+(renderHeight/2); y++ {
		fmt.Printf(" | ")
		for x := puzzle.guardPosition.x - (renderWidth / 2); x < puzzle.guardPosition.x+(renderWidth/2); x++ {
			output := ' '

			if y >= 0 && y < puzzle.boardHeight && x >= 0 && x < puzzle.boardWidth {
				output = puzzle.board[y][x]
			}

			if y == puzzle.guardPosition.y && x == puzzle.guardPosition.x {
				output = '^'
			}

			fmt.Printf("%c", output)
		}
		fmt.Printf(" | \n")
	}
	fmt.Printf("-----------------------------------\n")
}

func main() {
	puzzle := puzzleInstance{}

	boardData, err := os.ReadFile("6/input.txt")

	if err != nil {
		panic(err)
	}

	for y, line := range strings.Split(string(boardData), "\n") {
		puzzle.board = append(puzzle.board, []rune{})

		for x, r := range line {
			if r == '#' {
				puzzle.board[y] = append(puzzle.board[y], '#')
			} else if r == '^' {
				puzzle.guardVector = 0
				puzzle.guardPosition = position{
					x: x,
					y: y,
				}
				puzzle.board[y] = append(puzzle.board[y], 'X')
			} else {
				puzzle.board[y] = append(puzzle.board[y], ' ')
			}
		}
	}

	puzzle.boardHeight = len(puzzle.board)
	puzzle.boardWidth = len(puzzle.board[0])

	//fmt.Printf("Total: %d\n", part1Run(&puzzle))
	fmt.Printf("Total: %d\n", part2Run(&puzzle))
}

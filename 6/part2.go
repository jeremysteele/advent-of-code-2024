package main

import (
	"fmt"
)

func part2Run(puzzle *puzzleInstance) int {
	startX := puzzle.guardPosition.x
	startY := puzzle.guardPosition.y
	startVector := puzzle.guardVector

	obstacles := []string{}

	for y, row := range puzzle.board {
		for x, r := range row {
			visited := make(map[string]bool)
			visited[fmt.Sprintf("%d-%d-%d", startX, startY, startVector)] = true
			puzzle.guardVector = startVector
			puzzle.guardPosition.x = startX
			puzzle.guardPosition.y = startY

			if r != '#' && !(y == startY && x == startX) {
				puzzle.board[y][x] = '#'

				for puzzle.isPositionInbounds(puzzle.guardPosition) {
					isColliding := false

					for !isColliding {
						guardVector := directions[puzzle.guardVector]

						newPos := puzzle.guardPosition
						newPos.x += guardVector.x
						newPos.y += guardVector.y

						if !puzzle.isPositionInbounds(newPos) {
							puzzle.guardPosition = newPos
							break
						}

						key := fmt.Sprintf("%d-%d-%d", newPos.x, newPos.y, puzzle.guardVector)

						if _, ok := visited[key]; ok {
							obstacles = append(obstacles, key)
							puzzle.guardPosition.y = -1
							puzzle.guardPosition.x = -1
							break
						}

						visited[key] = true

						if puzzle.board[newPos.y][newPos.x] == '#' {
							isColliding = true
							puzzle.guardVector++

							if puzzle.guardVector >= maxDirection {
								puzzle.guardVector = 0
							}
							break
						} else {
							puzzle.board[newPos.y][newPos.x] = 'X'

							puzzle.guardPosition = newPos

							//puzzle.render()
							//time.Sleep(20 * time.Millisecond)
						}
					}
				}

				puzzle.board[y][x] = ' '
			}
		}
	}

	return len(obstacles)
}

package main

import "fmt"

func part1Run(puzzle *puzzleInstance) int {
	visited := make(map[string]bool)
	visited[fmt.Sprintf("%d-%d", puzzle.guardPosition.x, puzzle.guardPosition.y)] = true

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

			if puzzle.board[newPos.y][newPos.x] == '#' {
				isColliding = true
				puzzle.guardVector++

				if puzzle.guardVector >= maxDirection {
					puzzle.guardVector = 0
				}
				break
			} else {
				puzzle.board[newPos.y][newPos.x] = 'X'
				visited[fmt.Sprintf("%d-%d", newPos.x, newPos.y)] = true
				puzzle.guardPosition = newPos

				//puzzle.render()
				//time.Sleep(50 * time.Millisecond)
			}
		}
	}

	return len(visited)
}

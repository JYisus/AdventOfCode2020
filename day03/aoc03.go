package day03

import (
	"strings"
)

type Step struct {
	right int
	above int
}

type Position struct {
	column int
	row    int
}

func CountTreesInSlope(mapOfTrees string) (int, error) {
	grid := strings.Split(mapOfTrees, "\n")

	gridRows := len(grid)
	gridColumns := len(grid[0])

	step := Step{right: 3, above: 1}

	actualPosition := Position{column: 0, row: 0}

	trees := 0

	for actualPosition.row < gridRows {

		if string(grid[actualPosition.row][actualPosition.column%gridColumns]) == "#" {
			trees++
		}
		actualPosition = Position{
			column: actualPosition.column + step.right,
			row:    actualPosition.row + step.above,
		}
	}

	return trees, nil
}

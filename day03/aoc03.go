package day03

import (
	"strings"
)

type Step struct {
	right int
	down  int
}

type Position struct {
	column int
	row    int
}

func AoC03(mapOfTrees string, slopes []Step) (int, error) {

	result := 1

	for _, step := range slopes {
		trees, _ := CountTreesInSlope(mapOfTrees, step)
		result *= trees
	}

	return result, nil
}

func CountTreesInSlope(mapOfTrees string, step Step) (int, error) {
	grid := strings.Split(mapOfTrees, "\n")

	gridRows := len(grid)
	gridColumns := len(grid[0])

	actualPosition := Position{column: 0, row: 0}

	trees := 0

	for actualPosition.row < gridRows {

		if string(grid[actualPosition.row][actualPosition.column%gridColumns]) == "#" {
			trees++
		}
		actualPosition = Position{
			column: actualPosition.column + step.right,
			row:    actualPosition.row + step.down,
		}
	}

	return trees, nil
}

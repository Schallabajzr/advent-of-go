package day7

import (
	"fmt"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        7,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	grid := parseInput(input)
	result := 0
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ { // row
			if grid[i-1][j] != '|' && grid[i-1][j] != 'S' {
				continue
			}

			if grid[i][j] == '^' {
				if j > 0 {
					grid[i][j-1] = '|'
				}
				if j < len(grid[i])-1 {
					grid[i][j+1] = '|'
				}
				result++
			} else {
				grid[i][j] = '|'
			}
		}
		// printMatrixWide(grid)
	}

	return fmt.Sprintf("%d", result), nil
}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func printMatrixWide(matrix [][]rune) {
	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%2c", cell) // 2-char wide
		}
		fmt.Println()
	}
	fmt.Println()
}

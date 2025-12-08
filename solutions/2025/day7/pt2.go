package day7

import (
	"fmt"
	"strings"

	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        7,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	grid := parseInput(input)
	if len(grid) == 0 {
		return "0", nil
	}

	// counts grid: stores accumulated path counts
	counts := make([][]int, len(grid))
	for i := range grid {
		counts[i] = make([]int, len(grid[i]))
	}

	// locate S and set its initial path count to 1
	for i, row := range grid {
		for j, c := range row {
			if c == 'S' {
				counts[i][j] = 1
			}
		}
	}

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			aboveCount := counts[i-1][j]
			if aboveCount == 0 {
				continue
			}

			// splitter
			if grid[i][j] == '^' {
				if j > 0 {
					counts[i][j-1] += aboveCount
					grid[i][j-1] = '|'
				}
				if j < len(grid[i])-1 {
					counts[i][j+1] += aboveCount
					grid[i][j+1] = '|'
				}
				continue
			} else {
				counts[i][j] += aboveCount
			}
		}

		//printCountsWithSplits(grid, counts)
	}

	return fmt.Sprintf("%d", sumLastRow(counts)), nil
}

func sumLastRow(counts [][]int) int {
	sum := 0
	for _, v := range counts[len(counts)-1] {
		sum += v
	}
	return sum
}

func printCountsWithSplits(grid [][]rune, counts [][]int) {
	var b strings.Builder

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			// always show splitters
			if grid[i][j] == '^' {
				b.WriteString(" ^")
				continue
			}

			// show dot for empty cells
			if counts[i][j] == 0 {
				b.WriteString(" .")
				continue
			}

			// show number
			fmt.Fprintf(&b, "%2d", counts[i][j])
		}
		b.WriteByte('\n')
	}

	fmt.Print(b.String())
}

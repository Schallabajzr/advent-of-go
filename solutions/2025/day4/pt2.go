package day4

import (
	"fmt"
	"strings"

	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	dirs := []struct{ dx, dy int }{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	result := 0
	for {
		count := 0

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] != '@' {
					continue
				}

				neighborCount := 0
				for _, d := range dirs {
					nx, ny := x+d.dx, y+d.dy
					if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
						if grid[ny][nx] == '@' {
							neighborCount++
						}
					}
				}

				if neighborCount < 4 {
					count++
					grid[y][x] = 'x' // mark the cell as “changed”
				}
			}
		}

		result += count
		if count == 0 {
			// exit when no changes in this round
			break
		}
	}

	return fmt.Sprintf("%d", result), nil
}

func gridToFixedWidthString(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		for _, r := range row {
			sb.WriteString(fmt.Sprintf("%2c", r)) // 2-character width
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

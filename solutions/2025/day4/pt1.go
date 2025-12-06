package day4

import (
	"advent-of-go/utils"
	"fmt"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	count := 0
	dirs := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

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
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

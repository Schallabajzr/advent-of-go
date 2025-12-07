package day6

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        6,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < 2 {
		return "", fmt.Errorf("not enough lines")
	}

	// Last line contains operators
	opLine := strings.Fields(lines[len(lines)-1])
	numLines := lines[:len(lines)-1]

	cols := len(strings.Fields(numLines[0]))

	sums := make([]int, cols)
	prods := make([]int, cols)
	for i := range prods {
		prods[i] = 1
	}

	// Process numeric rows
	for _, l := range numLines {
		fields := strings.Fields(l)
		for col := 0; col < cols; col++ {
			n, err := strconv.Atoi(fields[col])
			if err != nil {
				return "", err
			}
			sums[col] += n
			prods[col] *= n
		}
	}

	// Combine results according to operator row
	total := 0
	for col := 0; col < cols; col++ {
		switch opLine[col] {
		case "+":
			total += sums[col]
		case "*":
			total += prods[col]
		default:
			return "", fmt.Errorf("invalid operator: %s", opLine[col])
		}
	}

	return fmt.Sprintf("%d", total), nil
}

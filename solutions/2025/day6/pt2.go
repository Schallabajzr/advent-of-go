package day6

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        6,
		Part:       2,
		Calculator: pt2,
	}
}

// rotateLeft90 rotates a slice of strings 90° to the left (character-wise)
func rotateLeft90(rows []string) []string {
	if len(rows) == 0 {
		return nil
	}

	// Find max row length
	maxLen := 0
	for _, row := range rows {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}

	// Pad rows with spaces to equal length
	for i, row := range rows {
		if len(row) < maxLen {
			rows[i] = row + strings.Repeat(" ", maxLen-len(row))
		}
	}

	// Build rotated output
	rotated := make([]string, maxLen)
	for col := maxLen - 1; col >= 0; col-- {
		var sb strings.Builder
		for row := 0; row < len(rows); row++ {
			sb.WriteByte(rows[row][col])
		}
		rotated[maxLen-1-col] = sb.String()
	}

	return rotated
}

func pt2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < 2 {
		return "", fmt.Errorf("not enough lines")
	}

	ops := strings.Fields(lines[len(lines)-1])
	rotated := rotateLeft90(lines[:len(lines)-1])

	var total int

	var res int
	opIndex := len(ops) - 1
	started := false

	for _, line := range rotated {
		line = strings.TrimSpace(line)
		if line == "" {
			// reset line
			opIndex--
			total += res
			res = 0
			started = false
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		if !started {
			res = num
			started = true
		} else if opIndex < len(ops) {
			switch ops[opIndex] {
			case "+":
				res += num
			case "*":
				res *= num
			default:
				return "", fmt.Errorf("unknown operator %s", ops[opIndex])
			}
		}
	}
	if started {
		total += res
	}
	return strconv.Itoa(total), nil
}

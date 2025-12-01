package day1

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	const (
		startPos = 50
		boundary = 100
	)

	position := startPos
	hits := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, s := range lines {
		if s == "" {
			continue
		}

		direction, steps, err := parseInstruction(s)
		steps %= boundary // full turns are discarded
		if err != nil {
			return "", err
		}

		switch direction {
		case "R":
			position += steps
			if position > boundary {
				position -= boundary
			}
		case "L":
			position -= steps
			if position < 0 {
				position += boundary
			}
		default:
			return "", fmt.Errorf("invalid direction: %s", direction)
		}

		if position == boundary {
			position = 0 // 100 is the same as 0
		}

		if position == 0 {
			hits++
		}
	}

	return fmt.Sprintf("%d", hits), nil
}

func parseInstruction(s string) (string, int, error) {
	if len(s) < 2 {
		return "", 0, fmt.Errorf("instruction too short: %s", s)
	}
	dist, err := strconv.Atoi(s[1:])
	if err != nil {
		return "", 0, err
	}
	return s[:1], dist, nil
}

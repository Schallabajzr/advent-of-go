package day1

import (
	"fmt"
	"strings"

	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
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

		dir, steps, err := parseInstruction(s)
		if err != nil {
			return "", err
		}

		fullRot := steps / boundary
		hits += fullRot
		steps = steps % boundary

		switch dir {
		case "R":
			prev := position
			position += steps
			if position > boundary {
				position -= boundary
				if prev%boundary != 0 {
					hits++
				}
			}

		case "L":
			prev := position
			position -= steps
			if position < 0 {
				position += boundary
				if prev%boundary != 0 {
					hits++
				}
			}

		default:
			return "", fmt.Errorf("invalid direction: %s", dir)
		}

		if position == boundary {
			position = 0
		}

		if position == 0 {
			hits++
		}
	}

	return fmt.Sprintf("%d", hits), nil
}

package day9

import (
	"errors"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        9,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "", errors.New("no input")
	}

	type point struct{ x, y int }
	points := make([]point, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), ",")
		if len(parts) != 2 {
			return "", errors.New("invalid coordinate: " + line)
		}
		x, err1 := strconv.Atoi(parts[0])
		y, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return "", errors.New("invalid number in: " + line)
		}
		points = append(points, point{x, y})
	}

	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			w := abs(p1.x-p2.x) + 1
			h := abs(p1.y-p2.y) + 1
			area := w * h

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return strconv.Itoa(maxArea), nil
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

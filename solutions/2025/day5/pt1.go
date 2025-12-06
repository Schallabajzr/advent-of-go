package day5

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        5,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format")
	}

	// Parse ranges
	rangeLines := strings.Split(parts[0], "\n")
	ranges := make([][2]int, 0, len(rangeLines))
	for _, line := range rangeLines {
		bounds := strings.Split(line, "-")
		if len(bounds) != 2 {
			return "", fmt.Errorf("invalid range: %s", line)
		}
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			return "", err
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			return "", err
		}
		ranges = append(ranges, [2]int{start, end})
	}

	// Parse IDs
	idLines := strings.Split(parts[1], "\n")
	ids := make([]int, 0, len(idLines))
	for _, line := range idLines {
		id, err := strconv.Atoi(line)
		if err != nil {
			return "", err
		}
		ids = append(ids, id)
	}

	// Count how many IDs fall inside any range
	count := 0
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count), nil
}

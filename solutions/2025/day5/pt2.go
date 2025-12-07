package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        5,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid input format")
	}

	// Parse ranges
	rangeLines := strings.Split(parts[0], "\n")
	ranges := make([][2]int, 0, len(rangeLines))
	for _, line := range rangeLines {
		bounds := strings.Split(line, "-")
		start, _ := strconv.Atoi(bounds[0])
		end, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, [2]int{start, end})
	}

	// Merge overlapping ranges
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	merged := make([][2]int, 0, len(ranges))
	for _, r := range ranges {
		if len(merged) == 0 || r[0] > merged[len(merged)-1][1] {
			merged = append(merged, r)
		} else {
			if r[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = r[1]
			}
		}
	}

	count := 0
	for _, r := range merged {
		count += r[1] - r[0] + 1
	}

	return strconv.Itoa(count), nil
}

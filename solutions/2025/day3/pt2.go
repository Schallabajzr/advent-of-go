package day3

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		if len(line) < 12 {
			continue
		}

		// Use runes to avoid string allocations inside the loop
		runes := []rune(line)
		tracker := make([]rune, 12)
		copy(tracker, runes[:12])

	NextChar:
		for _, c := range runes[12:] {
			for i := 0; i < 11; i++ {
				if tracker[i+1] > tracker[i] {
					copy(tracker[i:], tracker[i+1:])
					tracker[11] = c
					continue NextChar
				}
			}

			if c > tracker[11] {
				tracker[11] = c
			}
		}

		val, _ := strconv.Atoi(string(tracker))
		total += val
	}
	return fmt.Sprintf("%d", total), nil
}

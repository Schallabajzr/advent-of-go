package day2

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	parts := strings.Split(input, ",")
	count := 0
	for _, part := range parts {
		subParts := strings.Split(part, "-")
		left, _ := strconv.Atoi(subParts[0])
		right, _ := strconv.Atoi(subParts[1])

		for i := left; i <= right; i++ {
			if isRepeating(strconv.Itoa(i)) {
				count += i
			}
		}
	}
	return fmt.Sprintf("%d", count), nil
}

func isRepeating(s string) bool {
	if len(s) == 0 {
		return false
	}
	doubled := s + s
	middle := doubled[1 : len(doubled)-1]
	return strings.Contains(middle, s)
}

package day2

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	parts := strings.Split(input, ",")
	count := 0
	for _, part := range parts {
		subParts := strings.Split(part, "-")
		left, _ := strconv.Atoi(subParts[0])
		right, _ := strconv.Atoi(subParts[1])

		for i := left; i <= right; i++ {
			if isSymmetrical(i) {
				count += i
			}
		}
	}
	return fmt.Sprintf("%d", count), nil
}

func isSymmetrical(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

package day3

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

func pt1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}

		// Use runes to avoid string allocations inside the loop
		runes := []rune(line)
		first := runes[0]
		second := runes[1]

		for _, c := range runes[2:] {
			if second > first {
				first = second
				second = c
			} else if c > second {
				second = c
			}
		}

		// Use %c to format runes back to string representation
		num := fmt.Sprintf("%c%c", first, second)
		val, _ := strconv.Atoi(num)
		total += val
	}
	return fmt.Sprintf("%d", total), nil
}

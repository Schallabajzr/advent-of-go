package solutions

import (
	"slices"

	y2025 "advent-of-go/solutions/2025"
	"advent-of-go/utils"
)

func Solutions() []utils.Solution {
	return slices.Concat[[]utils.Solution](y2025.Solutions())
}

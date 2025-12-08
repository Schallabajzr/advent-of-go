package day8

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

func pt2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	pts := make([]Point, 0, len(lines))
	for _, l := range lines {
		f := strings.Split(l, ",")
		if len(f) != 3 {
			return "", errors.New("invalid input line")
		}
		x, _ := strconv.Atoi(strings.TrimSpace(f[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(f[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(f[2]))
		pts = append(pts, Point{int64(x), int64(y), int64(z)})
	}

	n := len(pts)

	// Build all edges (squared distance)
	edges := make([]Edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := pts[i].x - pts[j].x
			dy := pts[i].y - pts[j].y
			dz := pts[i].z - pts[j].z
			dist := dx*dx + dy*dy + dz*dz
			edges = append(edges, Edge{i, j, dist})
		}
	}

	// Sort edges ascending by distance
	sort.Slice(edges, func(i, j int) bool { return edges[i].d < edges[j].d })

	uf := NewUF(n)
	var lastI, lastJ int

	// Connect edges until all junction boxes are in one set
	for _, e := range edges {
		if uf.Union(e.i, e.j) {
			lastI, lastJ = e.i, e.j
			if len(uf.roots) == 1 {
				break
			}
		}
	}

	// Multiply the X coordinates of the last connected junction boxes
	return strconv.FormatInt(pts[lastI].x*pts[lastJ].x, 10), nil
}

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        8,
		Part:       2,
		Calculator: pt2,
	}
}

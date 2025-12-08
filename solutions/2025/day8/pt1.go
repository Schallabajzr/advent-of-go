package day8

import (
	"container/heap"
	"errors"
	"sort"
	"strconv"
	"strings"

	"advent-of-go/utils"
)

type (
	Point struct{ x, y, z int64 }
	Edge  struct {
		i, j int
		d    int64
	}
	MaxHeap []Edge
)

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i].d > h[j].d } // max-heap
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(Edge)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// UF (Union-Find) tracks disjoint sets and sizes, with root tracking.
type UF struct {
	p, s  []int
	roots map[int]bool
}

func NewUF(n int) *UF {
	uf := &UF{make([]int, n), make([]int, n), make(map[int]bool)}
	for i := 0; i < n; i++ {
		uf.p[i] = i
		uf.s[i] = 1
		uf.roots[i] = true
	}
	return uf
}

func (uf *UF) Find(a int) int {
	if uf.p[a] != a {
		uf.p[a] = uf.Find(uf.p[a])
	}
	return uf.p[a]
}

func (uf *UF) Union(a, b int) bool {
	ra, rb := uf.Find(a), uf.Find(b)
	if ra == rb {
		return false
	}
	if uf.s[ra] < uf.s[rb] {
		ra, rb = rb, ra
	}
	uf.p[rb] = ra
	uf.s[ra] += uf.s[rb]
	delete(uf.roots, rb)
	return true
}

// pt1K is a curried version that takes K as a parameter (edges to connect).
func pt1(K int) func(string) (string, error) {
	return func(input string) (string, error) {
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
		h := &MaxHeap{}
		heap.Init(h)

		// Build max-heap of K smallest edges
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				dx := pts[i].x - pts[j].x
				dy := pts[i].y - pts[j].y
				dz := pts[i].z - pts[j].z
				dist := dx*dx + dy*dy + dz*dz
				e := Edge{i, j, dist}
				if h.Len() < K {
					heap.Push(h, e)
				} else if dist < (*h)[0].d {
					heap.Pop(h)
					heap.Push(h, e)
				}
			}
		}

		// Extract edges and sort ascending
		edges := make([]Edge, 0, h.Len())
		for h.Len() > 0 {
			edges = append(edges, heap.Pop(h).(Edge))
		}
		sort.Slice(edges, func(i, j int) bool { return edges[i].d < edges[j].d })

		uf := NewUF(n)
		for _, e := range edges {
			uf.Union(e.i, e.j)
		}

		// Collect sizes of current roots
		sizes := make([]int, 0, len(uf.roots))
		for r := range uf.roots {
			sizes = append(sizes, uf.s[r])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

		if len(sizes) < 3 {
			return "0", nil
		}
		return strconv.Itoa(sizes[0] * sizes[1] * sizes[2]), nil
	}
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        8,
		Part:       1,
		Calculator: pt1(1000),
	}
}

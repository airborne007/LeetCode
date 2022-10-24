package pathwithminimumeffort

import (
	"sort"
)

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// Solution1: Binary search
func minimumEffortPath1(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}

		visited[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if p.x == m-1 && p.y == n-1 {
				return true
			}
			for _, dir := range dirs {
				x, y := p.x+dir.x, p.y+dir.y
				if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
					visited[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Solution2: Union Find
type UnionFind struct {
	root []int
	size []int
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	fx, fy := uf.Find(x), uf.Find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.root[fy] = fx
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func NewUnionFind(n int) *UnionFind {
	root := make([]int, n)
	size := make([]int, n)
	for i := range root {
		root[i] = i
		size[i] = 1
	}
	return &UnionFind{root, size}
}

type edge struct {
	v, w, diff int
}

func minimumEffortPath2(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	edges := []edge{}
	for i, row := range heights {
		for j, h := range row {
			id := i*n + j
			if i > 0 {
				edges = append(edges, edge{id - n, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })

	uf := NewUnionFind(m * n)
	for _, e := range edges {
		uf.Union(e.v, e.w)
		if uf.IsSame(0, n*m-1) {
			return e.diff
		}
	}
	return 0
}

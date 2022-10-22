package main

import "fmt"

// Solution1: Hash table
func longestConsecutive1(nums []int) int {
	ans := 0
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	for _, num := range nums {
		if m[num-1] {
			continue
		}
		cur := num
		length := 1
		for m[cur+1] {
			cur = cur + 1
			length++
		}
		if length > ans {
			ans = length
		}
	}
	return ans
}

// Solution2: Union-Find
type UnionFind struct {
	root map[int]int
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(x, y int) {
	xroot, yroot := uf.Find(x), uf.Find(y)
	if xroot == yroot {
		return
	}
	uf.root[xroot] = yroot
}
func (uf *UnionFind) Exist(x int) bool {
	_, ok := uf.root[x]
	return ok
}

func (uf *UnionFind) Count() int {
	ans := 0
	for k, v := range uf.root {
		ans = max(ans, uf.Find(v)-k+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func NewUnionFind(nums []int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make(map[int]int)
	for _, num := range nums {
		uf.root[num] = num
	}
	return uf
}

func longestConsecutive(nums []int) int {
	uf := NewUnionFind(nums)
	for _, num := range nums {
		if uf.Exist(num + 1) {
			uf.Union(num, num+1)
		}
	}
	return uf.Count()
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

package accountsmerge

import (
	"sort"
)

type UnionFind struct {
	root []int
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(from, to int) {
	uf.root[uf.Find(from)] = uf.Find(to)
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, n)
	for i := range uf.root {
		uf.root[i] = i
	}
	return uf
}

func accountsMerge(accounts [][]string) [][]string {
	emailToIndex := make(map[string]int)
	emailToName := make(map[string]string)
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, ok := emailToIndex[email]; !ok {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}

	uf := NewUnionFind(len(emailToIndex))
	for _, account := range accounts {
		firstEmail := account[1]
		for _, email := range account[2:] {
			uf.Union(emailToIndex[email], emailToIndex[firstEmail])
		}
	}

	indexToEmail := make(map[int][]string)
	for email, index := range emailToIndex {
		index = uf.Find(index)
		indexToEmail[index] = append(indexToEmail[index], email)
	}

	ans := make([][]string, 0)
	for _, emails := range indexToEmail {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}

	return ans
}

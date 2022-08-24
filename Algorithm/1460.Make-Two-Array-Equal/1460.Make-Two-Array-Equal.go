package maketwoarrayequal

import "sort"

// Solution1: use hashtab
func canBeEqual(target []int, arr []int) bool {
	cnt := make(map[int]int)
	for i, num := range target {
		cnt[num]++
		cnt[arr[i]]--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}

// Solution2: sort
func canBeEqual1(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i, c := range target {
		if arr[i] != c {
			return false
		}
	}
	return true
}

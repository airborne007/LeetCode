package findkclosestelements

import (
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Solution1: sort
func findClosestElements(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	ans := arr[:k]
	sort.Ints(ans)
	return ans
}

// Solution2: binary search + two pointer
func findClosestElements1(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right == len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}

package main

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums[left+1:], target+1) + left
	return []int{left, right}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3,4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1,-1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1,-1]
}

package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, n-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
			} else {
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
}

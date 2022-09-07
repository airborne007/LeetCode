package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-2 && nums[i]+nums[i+1]+nums[i+2] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		for left, right := i+1, n-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

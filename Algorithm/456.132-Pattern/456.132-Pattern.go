package pattern

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	stk := []int{nums[n-1]}
	maxK := math.MinInt
	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(stk) > 0 && nums[i] > stk[len(stk)-1] {
			maxK = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		if nums[i] > maxK {
			stk = append(stk, nums[i])
		}
	}

	return false
}

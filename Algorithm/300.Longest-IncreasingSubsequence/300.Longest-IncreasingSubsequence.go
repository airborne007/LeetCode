package longestincreasingsubsequence

import "sort"

// Solution1: Dynamic programe
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i, num := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < num {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Solution2: Greedy and binary search
func lengthOfLIS(nums []int) int {
	tails := make([]int, 0)
	for _, num := range nums {
		if len(tails) == 0 || num > tails[len(tails)-1] {
			tails = append(tails, num)
		} else {
			idx := sort.SearchInts(tails, num)
			tails[idx] = num
		}
	}
	return len(tails)
}

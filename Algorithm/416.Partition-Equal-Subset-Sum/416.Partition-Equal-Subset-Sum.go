package partitionequalsubsetsum

func canPartition(nums []int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}

	if total%2 == 1 {
		return false
	}

	target := total / 2
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

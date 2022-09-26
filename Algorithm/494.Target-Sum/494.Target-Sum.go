package targetsum

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	n := diff / 2
	dp := make([]int, n+1)
	dp[0] = 1
	for _, num := range nums {
		for j := n; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[n]
}

package integerbreak

// Solution1: DP
func integerBreak1(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], max(j*dp[i-j], j*(i-j)))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Solution2: greedy
func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		ans *= 3
		n -= 3
	}
	return ans * n
}

package perfectsquares

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans := math.MaxInt
		for j := 1; j*j <= i; j++ {
			ans = min(ans, dp[i-j*j])
		}
		dp[i] = ans + 1
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

package besttimetobuyandsellstockwithcooldown

// Solution1: DP
func maxProfit1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// dp[i][0]: the max profit of hold the stock
	// dp[i][1]: the max profit of don't hold the stock and in colddown
	// dp[i][2]: the max profit of don't hold the stock and not in colddown
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

// Solution2: DP with space optimization
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1, dp2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		f0 := max(dp0, dp2-prices[i])
		f1 := dp0 + prices[i]
		f2 := max(dp1, dp2)
		dp0, dp1, dp2 = f0, f1, f2
	}
	return max(dp1, dp2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

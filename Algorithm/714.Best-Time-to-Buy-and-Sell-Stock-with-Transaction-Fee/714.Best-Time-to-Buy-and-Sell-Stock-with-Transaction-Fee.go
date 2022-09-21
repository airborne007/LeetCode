package besttimetobuyandsellstockwithtransactionfee

// Solution1: DP
func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return dp[n-1][1]
}

// Solution2: DP with space optimization
func maxProfit2(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := -prices[0], 0
	for i := 1; i < n; i++ {
		f0 := max(dp0, dp1-prices[i])
		f1 := max(dp1, dp0+prices[i]-fee)
		dp0, dp1 = f0, f1
	}
	return dp1
}

// Solution3: Greedy
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := prices[0] + fee
	ans := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			ans += prices[i] - buy
			buy = prices[i]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

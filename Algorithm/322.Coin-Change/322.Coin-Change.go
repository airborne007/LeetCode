package coinchange

func coinChange(coins []int, amount int) int {
	MAX := amount + 1
	dp := make([]int, MAX)
	dp[0] = 0
	for i := 1; i < MAX; i++ {
		cost := MAX
		for _, c := range coins {
			if i-c >= 0 {
				cost = min(cost, dp[i-c]+1)
			}
		}
		dp[i] = cost
	}

	if dp[amount] == MAX {
		return -1
	} else {
		return dp[amount]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

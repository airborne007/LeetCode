package laststoneweightii

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, weight := range stones {
		sum += weight
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, weight := range stones {
		for j := target; j >= weight; j-- {
			dp[j] = max(dp[j], dp[j-weight]+weight)
		}
	}
	return sum - 2*dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package besttimetobuyandsellstockiv

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// buy[i]: the max profit of buy on the ith transaction
	buy := make([]int, k)
	// sell[i]: the max profit of sell on the ith transaction
	sell := make([]int, k)
	for i := range buy {
		buy[i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				buy[0] = max(buy[0], -prices[i])
				sell[0] = max(sell[0], buy[0]+prices[i])
			} else {
				buy[j] = max(buy[j], sell[j-1]-prices[i])
				sell[j] = max(sell[j], buy[j]+prices[i])
			}
		}
	}
	return sell[k-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

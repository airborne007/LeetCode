package besttimetobuyandsellstockiii

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < n; i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return max(sell1, sell2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

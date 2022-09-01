package finalpriceswithaspecialdiscountinashop

	ans := make([]int, len(prices))
	 := []int{}
	 i, price := range prices {
	r len(stk) > 0 && prices[stk[len(stk)-1]] >= price {
		= stk[len(stk)-1]
			[:len(stk)-1]
			= prices[idx] - price

		append(stk, i)

	 _, idx := range stk {
	s[idx] = prices[idx]

	urn ans



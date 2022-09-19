package maximumproductsubarray

func maxProduct(nums []int) int {
	MAX, MIN, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := MAX, MIN
		MAX = max(mx*nums[i], max(nums[i], nums[i]*mn))
		MIN = min(mn*nums[i], min(nums[i], nums[i]*mx))
		ans = max(ans, MAX)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

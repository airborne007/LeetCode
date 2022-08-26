package maximumproductoftwoelementsinanarray

func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if a > b {
		a, b = b, a
	}
	for _, num := range nums[2:] {
		if num > b {
			a, b = b, num
		} else if num > a {
			a = num
		}
	}
	return (a - 1) * (b - 1)
}

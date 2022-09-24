package jumpgame

func canJump(nums []int) bool {
	n, rightMost := len(nums), 0
	for i := 0; i < n; i++ {
		if i > rightMost {
			return false
		}
		rightMost = max(rightMost, nums[i]+i)
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package jumpgameii

func jump(nums []int) int {
	ans := 0
	n, rightMost := len(nums), 0
	end := 0
	for i := 0; i < n-1; i++ {
		rightMost = max(rightMost, i+nums[i])
		if i == end {
			ans++
			end = rightMost
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

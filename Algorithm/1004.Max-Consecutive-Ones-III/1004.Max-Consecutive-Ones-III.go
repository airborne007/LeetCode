package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	ans := 0
	left, cnt := 0, 0
	for right, num := range nums {
		if num == 0 {
			cnt++
		}
		if cnt > k {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

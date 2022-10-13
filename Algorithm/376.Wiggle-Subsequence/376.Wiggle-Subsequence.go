package wigglesubsequence

func wiggleMaxLength(nums []int) int {
	ans := 1
	prev, cur := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i+1] - nums[i]
		if (cur > 0 && prev <= 0) || (cur < 0 && prev >= 0) {
			ans++
			prev = cur
		}
	}
	return ans
}

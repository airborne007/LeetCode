package longestcontinuousincreasingsubsequence

func findLengthOfLCIS(nums []int) int {
	ans, count := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}

package subarrayproductlessthank

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans, prod := 0, 1
	left := 0
	for right := 0; right < len(nums); right++ {
		prod *= nums[right]
		for ; left <= right && prod >= k; left++ {
			prod /= nums[left]
		}
		ans += right - left + 1
	}
	return ans
}

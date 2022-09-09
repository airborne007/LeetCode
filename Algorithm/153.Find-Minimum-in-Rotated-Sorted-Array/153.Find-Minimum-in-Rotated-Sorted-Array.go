package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[n-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}

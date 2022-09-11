package maximumlengthofrepeatedsubarray

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	if m < n {
		return findLength(nums2, nums1)
	}
	dp := make([]int, n+1)
	ans := 0
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] {
				dp[j] = dp[j+1] + 1
			} else {
				dp[j] = 0
			}
			ans = max(ans, dp[j])
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

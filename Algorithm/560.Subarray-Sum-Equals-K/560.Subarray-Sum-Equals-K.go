package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	ans, preSum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		preSum += num
		if cnt, ok := m[preSum-k]; ok {
			ans += cnt
		}
		m[preSum] += 1
	}
	return ans
}

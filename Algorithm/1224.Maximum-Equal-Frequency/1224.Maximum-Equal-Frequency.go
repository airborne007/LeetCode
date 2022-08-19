package maxinumequalfrequency

func maxEqualFreq(nums []int) int {
	ans := 0

	cnt, freq := make(map[int]int), make(map[int]int)
	maxFreq := 0
	for idx, num := range nums {
		if freq[cnt[num]] > 0 {
			freq[cnt[num]]--
		}
		cnt[num]++
		maxFreq = max(maxFreq, cnt[num])
		freq[cnt[num]]++
		// End with current num, there is only three condition:
		// 1. All the number only appear 1 or 0 times
		// 2. All the number appear `maxFreq` or `maxFreq-1` times
		// 3. All the number appear `maxFreq` times, except only one num appear 1 times
		if maxFreq == 1 || freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == idx+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == idx+1 && freq[1] == 1 {
			ans = idx + 1
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

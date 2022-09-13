package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	left, maxCnt := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = max(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

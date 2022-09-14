package maximizetheconfusionofanexam

func maxConsecutiveAnswers(answerKey string, k int) int {
	ans := 0
	cnt, tcnt, fcnt := 0, 0, 0
	left, cnt := 0, 0
	for right, ch := range answerKey {
		if ch == 'T' {
			tcnt++
		}
		if ch == 'F' {
			fcnt++
		}
		cnt = max(tcnt, fcnt)
		if abs(right-left+1-cnt) > k {
			if answerKey[left] == 'T' {
				tcnt--
			}
			if answerKey[left] == 'F' {
				fcnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

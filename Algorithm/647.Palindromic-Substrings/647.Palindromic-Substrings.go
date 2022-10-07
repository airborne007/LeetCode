package palindromicsubstrings

// Solution1: DP
func countSubstrings1(s string) int {
	ans, n := 0, len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				ans += 1
				dp[i][j] = true
			}
		}
	}
	return ans
}

// Solution2: Two Pointer
func countSubstrings(s string) int {
	ans, n := 0, len(s)
	for i := 0; i < n; i++ {
		ans += extend(s, i, i, n)
		ans += extend(s, i, i+1, n)
	}
	return ans
}

func extend(s string, i, j, n int) int {
	cnt := 0
	for i >= 0 && j < n && s[i] == s[j] {
		cnt += 1
		i--
		j++
	}
	return cnt
}

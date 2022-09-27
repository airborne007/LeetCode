package issubsequence

// Solution1: two pointer
func isSubsequence1(s string, t string) bool {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == m
}

// Solution2: DP
func isSubsequence(s string, t string) bool {
	n1, n2 := len(s), len(t)
	dp := make([][]int, n1)
	for i := range dp {
		dp[i] = make([]int, n2)
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s[i-1] == t[j-2] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n1][n2] == n1
}

package distinctsubsequences

func numDistinct(s string, t string) int {
	n1, n2 := len(s), len(t)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n1][n2]
}

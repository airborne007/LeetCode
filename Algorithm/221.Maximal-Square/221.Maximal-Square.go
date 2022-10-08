package maximalsquare

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	maxSide := 0
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = 1
		}
	}
	for i := range dp[0] {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxSide = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

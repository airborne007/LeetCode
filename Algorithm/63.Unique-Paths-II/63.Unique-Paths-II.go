package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	block := false
	for i := range dp {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			block = true
		}
		if !block {
			dp[i][0] = 1
		}
	}
	block = false
	for i := range dp[0] {
		if dp[0][i] == 1 {
			block = true
		}
		if !block {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

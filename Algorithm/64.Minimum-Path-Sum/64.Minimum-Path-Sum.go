package minimumpathsum

func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dp[i][j] = grid[i][j]
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] += dp[i][j-1]
			} else if j == 0 {
				dp[i][j] += dp[i-1][j]
			} else {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1])
			}

		}
	}
	return dp[rows-1][cols-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

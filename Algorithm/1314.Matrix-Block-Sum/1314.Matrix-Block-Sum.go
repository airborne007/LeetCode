package matrixblocksum

func matrixBlockSum(mat [][]int, k int) [][]int {
	rows, cols := len(mat), len(mat[0])
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + mat[i][j]
		}
	}

	ans := make([][]int, rows)
	for i := range ans {
		ans[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r1, c1 := max(i-k, 0), max(j-k, 0)
			r2, c2 := min(i+k, rows-1), min(j+k, cols-1)
			ans[i][j] = dp[r2+1][c2+1] - dp[r1][c2+1] - dp[r2+1][c1] + dp[r1][c1]
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

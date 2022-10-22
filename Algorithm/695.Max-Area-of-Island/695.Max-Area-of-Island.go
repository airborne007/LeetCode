package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans, temp := 0, 0
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || grid[row][col] == 0 {
			return
		}
		grid[row][col] = 0
		temp++
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			temp = 0
			dfs(i, j)
			ans = max(ans, temp)
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

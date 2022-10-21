package surroundedregions

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || row >= m || col < 0 || col >= n || board[row][col] != 'O' {
			return
		}
		board[row][col] = '-'
		for _, dir := range dirs {
			dfs(row+dir[0], col+dir[1])
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				continue
			}
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case '-':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}

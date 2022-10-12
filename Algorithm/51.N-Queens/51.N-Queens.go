package nqueens

func solveNQueens(n int) [][]string {
	var ans [][]string
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := make(map[int]bool)
	diagonals1, diagonals2 := make(map[int]bool), make(map[int]bool)
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			board := genBoard(queens, n)
			ans = append(ans, board)
			return
		}
		for i := 0; i < n; i++ {
			d1, d2 := row-i, row+i
			if columns[i] || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[i], diagonals1[d1], diagonals2[d2] = true, true, true
			queens[row] = i
			dfs(row + 1)
			queens[row] = -1
			columns[i], diagonals1[d1], diagonals2[d2] = false, false, false
		}
	}
	dfs(0)
	return ans
}

func genBoard(queens []int, n int) []string {
	board := make([]string, 0, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

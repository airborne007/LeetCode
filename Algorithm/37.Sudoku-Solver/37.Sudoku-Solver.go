package sudokusolver

func solveSudoku(board [][]byte) {
	dfs(board, 0, 0)
}

func dfs(board [][]byte, row, col int) bool {
	if row == 8 && col == 9 {
		return true
	}

	if col == 9 {
		row++
		col = 0
	}

	if board[row][col] != '.' {
		return dfs(board, row, col+1)
	}

	for num := byte('1'); num <= byte('9'); num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if dfs(board, row, col+1) {
				return true
			}
			board[row][col] = '.'
		}
	}
	return false
}

func isValid(board [][]byte, row, col int, k byte) bool {
	for i := 0; i < 9; i++ {
		// Rows
		if board[row][i] == k {
			return false
		}
		// Columns
		if board[i][col] == k {
			return false
		}

		// 3*3 Grid
		startRow, startCol := (row/3)*3, (col/3)*3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == k {
					return false
				}
			}
		}
	}
	return true
}

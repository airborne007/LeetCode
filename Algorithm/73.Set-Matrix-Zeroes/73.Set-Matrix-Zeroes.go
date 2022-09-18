package setmatrixzeroes

func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for _, num := range matrix[0] {
		if num == 0 {
			row0 = true
		}
	}
	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for i := 0; i < cols; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}

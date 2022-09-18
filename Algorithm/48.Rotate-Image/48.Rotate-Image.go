package rotateimage

func rotate(matrix [][]int) {
	n := len(matrix)
	// up down swap
	up, down := 0, n-1
	for up < down {
		matrix[up], matrix[down] = matrix[down], matrix[up]
		up++
		down--
	}

	// swap [i, j] with [j, i]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

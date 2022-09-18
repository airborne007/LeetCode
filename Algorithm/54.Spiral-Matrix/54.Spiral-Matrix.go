package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	ans := make([]int, 0, rows*cols)
	left, right, top, bottom := 0, cols-1, 0, rows-1
	for left <= right && top <= bottom {
		for index := left; index <= right; index++ {
			ans = append(ans, matrix[top][index])
		}
		for index := top + 1; index <= bottom; index++ {
			ans = append(ans, matrix[index][right])
		}
		if left < right && top < bottom {
			for index := right - 1; index >= left; index-- {
				ans = append(ans, matrix[bottom][index])
			}
			for index := bottom - 1; index >= top+1; index-- {
				ans = append(ans, matrix[index][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}

package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := 0, n-1; i >= 0 && i < m && j >= 0 && j < n; {
		num := matrix[i][j]
		if num == target {
			return true
		} else if num < target {
			i++
		} else {
			j--
		}
	}
	return false
}

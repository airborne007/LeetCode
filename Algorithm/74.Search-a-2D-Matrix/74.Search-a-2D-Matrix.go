package main

import (
	"fmt"
	"sort"
)

// Solution1, search from the right corner, time O(m+n), space O(1)
func searchMatrix1(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	i, j := 0, cols-1
	for i >= 0 && i < rows && j >= 0 && j < cols {
		x := matrix[i][j]
		if x == target {
			return true
		} else if x < target {
			i++
		} else {
			j--
		}
	}
	return false
}

// Solution2: binary search, time O(logmn), space O(1)
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	row := sort.Search(rows, func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < cols && matrix[row][col] == target
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}

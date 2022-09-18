package main

import "fmt"

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for index := left; index <= right; index++ {
			matrix[top][index] = num
			num++
		}
		for index := top + 1; index <= bottom; index++ {
			matrix[index][right] = num
			num++
		}
		if left < right && top < bottom {
			for index := right - 1; index >= left; index-- {
				matrix[bottom][index] = num
				num++
			}
			for index := bottom - 1; index >= top+1; index-- {
				matrix[index][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
}

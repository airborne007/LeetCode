package main

import "fmt"

// Solution1: time O(m*n), space O(m+n)
func numSpecial1(mat [][]int) int {
	ans := 0
	rowSum, colSum := make([]int, len(mat)), make([]int, len(mat[0]))
	for i, row := range mat {
		for j, x := range row {
			rowSum[i] += x
			colSum[j] += x
		}
	}

	for i, row := range mat {
		for j, x := range row {
			if x == 1 && rowSum[i] == 1 && colSum[j] == 1 {
				ans++
			}
		}
	}

	return ans
}

// Solution2: time O(m*n), space O(1)
func numSpecial(mat [][]int) (ans int) {
	for i, row := range mat {
		cnt1 := 0
		for _, x := range row {
			cnt1 += x
		}
		if i == 0 {
			cnt1--
		}
		if cnt1 > 0 {
			for j, x := range row {
				if x == 1 {
					mat[0][j] += cnt1
				}
			}
		}
	}
	for _, x := range mat[0] {
		if x == 1 {
			ans++
		}
	}
	return
}

func main() {
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}})) // output: 1
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}})) // output: 3
}

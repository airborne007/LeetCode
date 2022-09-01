package main

import "fmt"

func longestWPI(hours []int) int {
	n := len(hours)
	scores := make([]int, n)
	for i, hour := range hours {
		if hour > 8 {
			scores[i] = 1
		} else {
			scores[i] = -1
		}
	}

	prefixSum := make([]int, n+1)
	for i, score := range scores {
		prefixSum[i+1] = score + prefixSum[i]
	}

	stk := []int{}
	for i, sum := range prefixSum {
		if len(stk) == 0 || prefixSum[stk[len(stk)-1]] > sum {
			stk = append(stk, i)
		}
	}

	ans := 0
	for i := n; i >= 0; i-- {
		for len(stk) > 0 && prefixSum[i] > prefixSum[stk[len(stk)-1]] {
			ans = max(ans, i-stk[len(stk)-1])
			stk = stk[:len(stk)-1]
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 9}))
}

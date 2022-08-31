package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	idx := 0
	stk := []int{}
	for _, num := range pushed {
		stk = append(stk, num)
		for len(stk) > 0 && stk[len(stk)-1] == popped[idx] {
			stk = stk[:len(stk)-1]
			idx++
		}
	}
	return len(stk) == 0
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}

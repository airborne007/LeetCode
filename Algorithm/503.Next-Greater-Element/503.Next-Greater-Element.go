package nextgreaterelement

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	stk := make([]int, 0)
	for i := 0; i < 2*n-1; i++ {
		num := nums[i%n]
		for len(stk) > 0 && num > nums[stk[len(stk)-1]] {
			ans[stk[len(stk)-1]] = num
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, i%n)
	}
	return ans
}

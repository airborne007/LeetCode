package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		} else {
			ans[i] = 0
		}
		stack = append(stack, i)
	}
	return ans
}

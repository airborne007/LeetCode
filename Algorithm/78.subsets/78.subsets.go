package subsets

// Solution1： bit manipulation
func subsets1(nums []int) [][]int {
	ans := make([][]int, 0)
	for mask := 0; mask < 1<<len(nums); mask++ {
		res := make([]int, 0)
		for i, v := range nums {
			if mask>>i&1 > 0 {
				res = append(res, v)
			}
		}
		ans = append(ans, append([]int{}, res...))
	}
	return ans
}

// Solution2: Backtracking
func subsets2(nums []int) [][]int {
	ans := make([][]int, 0)
	res := make([]int, 0)
	n := len(nums)
	var dfs func(int)
	dfs = func(cur int) {
		if cur == n {
			ans = append(ans, append([]int{}, res...))
			return
		}
		res = append(res, nums[cur])
		dfs(cur + 1)
		res = res[:len(res)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

// Solution3: Dynamic programming
func subsets(nums []int) [][]int {
	dp := [][]int{{}}
	for _, num := range nums {
		for _, prev := range dp {
			dp = append(dp, append(append([]int{}, prev...), num))
		}
	}
	return dp
}

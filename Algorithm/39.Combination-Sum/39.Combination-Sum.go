package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	var sum int
	var dfs func(int)
	dfs = func(start int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(i)
			sum -= path[len(path)-1]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

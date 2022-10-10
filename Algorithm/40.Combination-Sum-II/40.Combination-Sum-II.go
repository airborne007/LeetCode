package combinationsumii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		ans  [][]int
		path []int
		sum  int
		dfs  func(int)
	)
	n := len(candidates)
	sort.Ints(candidates)
	dfs = func(start int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < n; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			dfs(i + 1)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

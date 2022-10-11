package subsetsii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(int)
	n := len(nums)
	sort.Ints(nums)
	dfs = func(start int) {
		ans = append(ans, append([]int{}, path...))
		for i := start; i < n; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

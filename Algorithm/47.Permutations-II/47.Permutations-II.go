package permutationsii

func permuteUnique(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
		dfs  func()
	)
	n := len(nums)
	used := make(map[int]bool)
	dfs = func() {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		seen := make(map[int]bool)
		for i := 0; i < n; i++ {
			if used[i] || seen[nums[i]] {
				continue
			}
			seen[nums[i]] = true
			used[i] = true
			path = append(path, nums[i])
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return ans
}

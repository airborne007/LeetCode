package permutations

func permute(nums []int) [][]int {
	var ans [][]int
	var path []int
	var used map[int]bool
	var dfs func()
	n := len(nums)
	dfs = func() {
		if len(path) == n {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
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

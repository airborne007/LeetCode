package increasingsubsequences

func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(int)
	n := len(nums)
	dfs = func(start int) {
		if len(path) >= 2 {
			ans = append(ans, append([]int{}, path...))
		}
		seen := make(map[int]bool)
		for i := start; i < n; i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || seen[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			seen[nums[i]] = true
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

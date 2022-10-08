package combinations

func combine(n int, k int) [][]int {
	var ans [][]int
	var temp []int
	var dfs func(int)
	dfs = func(start int) {
		if len(temp) == k {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		for i := start; i <= n-(k-len(temp))+1; i++ {
			temp = append(temp, i)
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(1)
	return ans
}

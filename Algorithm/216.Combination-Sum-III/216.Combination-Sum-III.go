package combinationsumiii

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var temp []int
	var sum int
	var dfs func(int)
	dfs = func(start int) {
		if sum > n {
			return
		}
		if len(temp) == k {
			if sum == n {
				ans = append(ans, append([]int{}, temp...))
				return
			}
			return
		}
		for i := start; i <= 9-(k-len(temp))+1; i++ {
			temp = append(temp, i)
			sum += i
			dfs(i + 1)
			sum -= temp[len(temp)-1]
			temp = temp[:len(temp)-1]
		}
	}
	dfs(1)
	return ans
}

package palindromepartitioning

func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var dfs func(string, int)
	dfs = func(s string, start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
		}
		for i := start; i < len(s); i++ {
			if !isPalindrome(s[start : i+1]) {
				continue
			}
			path = append(path, s[start:i+1])
			dfs(s, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(s, 0)
	return res
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	for left, right := 0, len(s)-1; left < right; {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

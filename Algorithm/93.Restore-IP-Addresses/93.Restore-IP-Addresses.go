package restoreipaddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	var path []string
	var dfs func(int)
	n := len(s)
	if n > 12 {
		return nil
	}
	dfs = func(start int) {
		if start == n && len(path) == 4 {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		if len(path) > 4 {
			return
		}
		for i := start; i < n; i++ {
			if i-start+1 <= 3 && len(path) <= 4 && isValid(s[start:i+1]) {
				path = append(path, s[start:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			} else {
				return
			}
		}
	}
	dfs(0)
	return ans
}

func isValid(str string) bool {
	if str == "" || (len(str) > 1 && str[0] == '0') {
		return false
	}
	num, _ := strconv.Atoi(str)
	return num <= 255 && num >= 0
}

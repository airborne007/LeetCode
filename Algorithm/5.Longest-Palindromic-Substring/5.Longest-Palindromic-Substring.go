package longestpalindromicsubstring

// Solution1: DP
func longestPalindrome1(s string) string {
	n := len(s)
	maxLen := 0
	left, right := 0, 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (dp[i+1][j-1] || j-i <= 1) {
				dp[i][j] = true
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				left = i
				right = j
			}
		}
	}
	return s[left:right]
}

// Solution2: Two Pointer
func longestPalindrome(s string) string {
	n := len(s)
	maxLen, left, right := 0, 0, 0

	var extend func(string, int, int, int)
	extend = func(s string, i, j, n int) {
		for i >= 0 && j < n && s[i] == s[j] {
			if j-i+1 > maxLen {
				maxLen = j - i + 1
				left = i
				right = j
			}
			i--
			j++
		}
	}
	for i := n - 1; i >= 0; i-- {
		extend(s, i, i, n)
		extend(s, i, i+1, n)
	}
	return s[left : right+1]
}

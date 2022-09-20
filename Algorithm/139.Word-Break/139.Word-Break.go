package wordbreak

func wordBreak(s string, wordDict []string) bool {
	wordset := make(map[string]bool)
	for _, word := range wordDict {
		wordset[word] = true
	}
	// dp[i]: s[0, i] is valid or not.
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordset[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

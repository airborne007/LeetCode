package lettercombinationsofaphonenumber

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	letters := [10]string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}

	n := len(digits)
	var ans []string
	var temp string
	var dfs func(int)
	dfs = func(start int) {
		if len(temp) == n {
			ans = append(ans, temp)
			return
		}
		word := letters[digits[start]-'0']
		for i := 0; i < len(word); i++ {
			temp += string(word[i])
			dfs(start + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return ans
}

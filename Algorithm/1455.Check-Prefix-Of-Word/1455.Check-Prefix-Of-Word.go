package checkprefixofword

func isPrefixOfWord(sentence string, searchWord string) int {
	ans := 1
	n := len(searchWord)
	i := 0
	for i < len(sentence) {
		if sentence[i] == ' ' {
			ans++
			i++
		}
		j := 0
		for ; j < n; j++ {
			if sentence[i] != searchWord[j] {
				break
			}
			i++
		}
		if j == n {
			return ans
		} else {
			// Skip to the next word
			for i < len(sentence) && sentence[i] != ' ' {
				i++
			}
		}
	}
	return -1
}

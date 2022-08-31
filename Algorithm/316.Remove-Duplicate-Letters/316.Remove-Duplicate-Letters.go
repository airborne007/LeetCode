package removeduplicateletters

func removeDuplicateLetters(s string) string {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	stk := []rune{}
	visited := [26]bool{}
	for _, ch := range s {
		if !visited[ch-'a'] {
			for len(stk) > 0 && stk[len(stk)-1] > ch {
				c := stk[len(stk)-1] - 'a'
				if cnt[c] == 0 {
					break
				}
				visited[c] = false
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, ch)
			visited[ch-'a'] = true
		}
		cnt[ch-'a']--
	}

	return string(stk)
}

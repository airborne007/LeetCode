package removekdigits

import "strings"

func removeKdigits(num string, k int) string {
	stk := []rune{}
	for _, ch := range num {
		for k > 0 && len(stk) > 0 && stk[len(stk)-1] > ch {
			stk = stk[:len(stk)-1]
			k--
		}
		stk = append(stk, ch)
	}
	stk = stk[:len(stk)-k]
	ans := strings.TrimLeft(string(stk), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

package minimumremovetomakevalidparentheses

import "strings"

func minRemoveToMakeValid1(s string) string {
	toRemoved := map[int]struct{}{}
	stk := []int{}
	for i, ch := range s {
		switch ch {
		case '(':
			stk = append(stk, i)
		case ')':
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			} else {
				toRemoved[i] = struct{}{}
			}
		}
	}

	// The remained '(' should be removed.
	for _, idx := range stk {
		toRemoved[idx] = struct{}{}
	}

	sb := &strings.Builder{}
	for i, ch := range s {
		if _, ok := toRemoved[i]; !ok {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

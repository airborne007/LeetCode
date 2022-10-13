package assigncookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			ans++
			i++
			j++
		} else {
			j++
		}
	}
	return ans
}

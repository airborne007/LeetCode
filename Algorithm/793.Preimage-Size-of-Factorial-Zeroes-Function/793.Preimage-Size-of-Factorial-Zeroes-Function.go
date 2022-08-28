package preimagesizeoffactorialzeroesfunction

import "sort"

func zeta(n int) int {
	ans := 0
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}

func nx(k int) int {
	return sort.Search(5*k, func(x int) bool { return zeta(x) >= k })
}

func preimageSizeFZF(k int) int {
	return nx(k+1) - nx(k)
}

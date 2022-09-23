package uglynumber

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range []int{2, 3, 5} {
		if n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

package pow

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	ans := 1.0
	contribute := x
	for n > 0 {
		if n&1 == 1 {
			ans *= contribute
		}
		contribute *= contribute
		n >>= 1
	}
	return ans
}

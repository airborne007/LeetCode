package singlenumberiii

// Solution1: Hash table
func singleNumber1(nums []int) []int {
	ans := make([]int, 0, 2)
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for num, cnt := range freq {
		if cnt == 1 {
			ans = append(ans, num)
		}
	}
	return ans
}

// Solution2: Bits manipulation
func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	a, b := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

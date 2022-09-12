package rabbitsinforest

func numRabbits(answers []int) int {
	cnt := make(map[int]int)
	for _, answer := range answers {
		cnt[answer]++
	}

	ans := 0
	for y, x := range cnt {
		ans += (x + y) / (y + 1) * (y + 1)
	}
	return ans
}

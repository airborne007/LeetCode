package getequalsubstringswithinbudget

func equalSubstring(s string, t string, maxCost int) int {
	left, cost := 0, 0
	for right := range s {
		cost += abs(int(s[right]) - int(t[right]))
		if cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}
	}
	return len(s) - left
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

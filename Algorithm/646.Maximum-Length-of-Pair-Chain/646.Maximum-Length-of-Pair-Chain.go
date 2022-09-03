package maximumlengthofpairchain

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	ans := 1
	prev := pairs[0][1]
	for _, pair := range pairs[1:] {
		if prev < pair[0] {
			ans++
			prev = pair[1]
		}
	}
	return ans
}

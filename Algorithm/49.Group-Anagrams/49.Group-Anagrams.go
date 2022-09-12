package groupanagrams

func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, ch := range str {
			cnt[ch-'a']++
		}
		group[cnt] = append(group[cnt], str)
	}
	ans := make([][]string, 0, len(group))
	for _, v := range group {
		ans = append(ans, v)
	}
	return ans
}

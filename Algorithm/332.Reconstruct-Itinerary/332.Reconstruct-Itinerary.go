package reconstructitinerary

import "sort"

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string)
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	for k := range m {
		sort.Strings(m[k])
	}

	ans := make([]string, 0)

	var dfs func(string)
	dfs = func(start string) {
		for {
			if v, ok := m[start]; !ok || len(v) == 0 {
				break
			}
			tmp := m[start][0]
			m[start] = m[start][1:]
			dfs(tmp)
		}
		ans = append(ans, start)
	}
	dfs("JFK")

	for left, right := 0, len(ans)-1; left < right; {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	return ans
}

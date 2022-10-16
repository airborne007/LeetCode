package coursescheduleii

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		ans      []int
		indegree = make([]int, numCourses)
		edges    = make([][]int, numCourses)
		queue    []int
	)

	for _, edge := range prerequisites {
		cur, pre := edge[0], edge[1]
		indegree[cur]++
		edges[pre] = append(edges[pre], cur)
	}

	for i, cnt := range indegree {
		if cnt == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		ans = append(ans, pre)
		for _, cur := range edges[pre] {
			indegree[cur]--
			if indegree[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	if len(ans) == numCourses {
		return ans
	} else {
		return nil
	}
}

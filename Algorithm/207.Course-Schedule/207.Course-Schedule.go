package courseschedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	edges := make([][]int, numCourses)

	for i := range prerequisites {
		cur, pre := prerequisites[i][0], prerequisites[i][1]
		indegree[cur]++
		edges[pre] = append(edges[pre], cur)
	}

	queue := make([]int, 0)
	for i, cnt := range indegree {
		if cnt == 0 {
			queue = append(queue, i)
		}
	}

	cnt := 0
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		cnt++
		for _, cur := range edges[pre] {
			indegree[cur]--
			if indegree[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	return cnt == numCourses
}

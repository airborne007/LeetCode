package numberofprovinces

// Solution1: DFS
func findCircleNum1(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	var dfs func(int)
	dfs = func(from int) {
		visited[from] = true
		for to, connected := range isConnected[from] {
			if connected == 1 && !visited[to] {
				dfs(to)
			}
		}
	}

	ans := 0
	for i, v := range visited {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans
}

// Solution2: BFS
func findCircleNum2(isConnected [][]int) int {
	ans := 0
	visited := make([]bool, len(isConnected))
	for i, v := range visited {
		if !v {
			ans++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				visited[from] = true
				for to, connected := range isConnected[from] {
					if connected == 1 && !visited[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}
	return ans
}

// Solution3: Union-Find
type UnionFind struct {
	root []int
}

func (uf *UnionFind) Find(x int) int {
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(from, to int) {
	uf.root[uf.Find(from)] = uf.Find(to)
}

func (uf *UnionFind) Count() int {
	cnt := 0
	for i, v := range uf.root {
		if i == v {
			cnt++
		}
	}
	return cnt
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.root = make([]int, n)
	for i := range uf.root {
		uf.root[i] = i
	}
	return uf
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(j, i)
			}
		}
	}
	return uf.Count()
}

package numberofislands

// Solution1: DFS
func numIslands1(grid [][]byte) int {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(row, col int) {
		grid[row][col] = '0'
		for _, dir := range dirs {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == '1' {
				dfs(newRow, newCol)
			}
		}
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

// Solution2: BFS
func numIslands2(grid [][]byte) int {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, cols := len(grid), len(grid[0])
	var bfs func(int, int)
	bfs = func(row, col int) {
		queue := [][2]int{{row, col}}
		for len(queue) > 0 {
			row, col = queue[0][0], queue[0][1]
			queue = queue[1:]
			if row >= 0 && row < rows && col >= 0 && col < cols && grid[row][col] == '1' {
				grid[row][col] = '0'
				for _, dir := range dirs {
					queue = append(queue, [2]int{row + dir[0], col + dir[1]})
				}
			}
		}
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				bfs(i, j)
				ans++
			}
		}
	}
	return ans
}

// Solution3: Union-Find
type UnionFind struct {
	cnt        int
	root, rank []int
}

func (uf *UnionFind) Find(x int) int {
	if x > len(uf.root) {
		return -1
	}
	if uf.root[x] != x {
		uf.root[x] = uf.Find(uf.root[x])
	}
	return uf.root[x]
}

func (uf *UnionFind) Union(p, q int) {
	pRoot, qRoot := uf.Find(p), uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.root[pRoot] = qRoot
	} else {
		uf.root[qRoot] = pRoot
	}
	if uf.rank[pRoot] == uf.rank[qRoot] {
		uf.rank[pRoot]++
	}
	uf.cnt--
}

func (u *UnionFind) Count() int {
	return u.cnt
}

func NewUnionFind(grid [][]byte) *UnionFind {
	m, n := len(grid), len(grid[0])
	uf := new(UnionFind)
	uf.root = make([]int, m*n)
	uf.rank = make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.cnt++
			}
			uf.root[i*n+j] = i*n + j
		}
	}
	return uf
}

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	uf := NewUnionFind(grid)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				if i+1 < rows && grid[i+1][j] == '1' {
					uf.Union(i*cols+j, (i+1)*cols+j)
				}
				if j+1 < cols && grid[i][j+1] == '1' {
					uf.Union(i*cols+j, i*cols+j+1)
				}
			}
		}
	}
	return uf.Count()
}

package shortestbridge

func shortestBridge(grid [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 0
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, num := range row {
			if num != 1 {
				continue
			}
			grid[i][j] = -1
			islands := []pair{}
			queue := []pair{{i, j}}
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				islands = append(islands, p)
				for _, d := range dirs {
					x, y := p.x+d.x, p.y+d.y
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
						grid[x][y] = -1
						queue = append(queue, pair{x, y})
					}
				}
			}

			queue = islands
			for {
				tmp := queue
				queue = nil
				for _, p := range tmp {
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if x >= 0 && x < m && y >= 0 && y < n {
							if grid[x][y] == 1 {
								return ans
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1
								queue = append(queue, pair{x, y})
							}
						}
					}
				}
				ans++
			}
		}
	}
	return ans
}

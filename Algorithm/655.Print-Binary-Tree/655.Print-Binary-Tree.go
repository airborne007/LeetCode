package printbinarytree

import (
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func printTree(root *TreeNode) [][]string {
	height := calTreeHeight(root)
	rows := height + 1
	cols := 1<<rows - 1
	ans := make([][]string, rows)
	for i := 0; i <= height; i++ {
		ans[i] = make([]string, cols)
	}
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(height-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(height-r-1))
		}
	}
	dfs(root, 0, (cols-1)/2)
	return ans
}

// BFS
func printTree1(root *TreeNode) [][]string {
	height := calTreeHeight(root)
	rows := height + 1
	cols := 1<<rows - 1
	ans := make([][]string, rows)
	for i := 0; i <= height; i++ {
		ans[i] = make([]string, cols)
	}
	if root == nil {
		return ans
	}
	type entry struct {
		node *TreeNode
		r, c int
	}
	q := []entry{{root, 0, (cols - 1) / 2}}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		node, r, c := e.node, e.r, e.c
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			q = append(q, entry{node.Left, r + 1, c - 1<<(height-r-1)})
		}
		if node.Right != nil {
			q = append(q, entry{node.Right, r + 1, c + 1<<(height-r-1)})
		}
	}

	return ans
}

func calTreeHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + max(calTreeHeight(root.Left), calTreeHeight(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

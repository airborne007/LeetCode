package findduplicatesubtrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	type pair struct {
		node  *TreeNode
		index int
	}
	repeated := map[*TreeNode]struct{}{}
	seen := map[[3]int]pair{}
	idx := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
		if p, ok := seen[tri]; ok {
			repeated[p.node] = struct{}{}
			return p.index
		}
		idx++
		seen[tri] = pair{node, idx}
		return idx
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeated))
	for node := range repeated {
		ans = append(ans, node)
	}
	return ans
}

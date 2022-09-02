package longestunivaluepath

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftPath := dfs(node.Left)
		rightPath := dfs(node.Right)
		var left, right int
		if node.Left != nil && node.Left.Val == node.Val {
			left = leftPath + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right = rightPath + 1
		}
		ans = max(ans, left+right)
		return max(left, right)
	}
	dfs(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

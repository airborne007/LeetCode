package houserobberiii

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	ans := helper(root)
	return max(ans[0], ans[1])
}

func helper(node *TreeNode) [2]int {
	// result[0]: don't rob, result[1]: rob
	var result [2]int
	if node == nil {
		return result
	}
	left := helper(node.Left)
	right := helper(node.Right)
	result[0] = max(left[0], left[1]) + max(right[0], right[1])
	result[1] = node.Val + left[0] + right[0]
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

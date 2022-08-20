package maximumbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	for i, num := range nums {
		if num > nums[idx] {
			idx = i
		}
	}
	root := &TreeNode{Val: nums[idx]}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}

// Monotonic stack
func constructMaximumBinaryTree1(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	stk := []*TreeNode{}
	for _, num := range nums {
		node := &TreeNode{Val: num}
		for len(stk) > 0 {
			top := stk[len(stk)-1]
			if num < top.Val {
				top.Right = node
				break
			} else {
				node.Left = top
				stk = stk[:len(stk)-1]
			}
		}
		stk = append(stk, node)
	}
	return stk[0]
}

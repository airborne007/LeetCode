package maximumwidthofbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Solution1: DFS
func widthOfBinaryTree(root *TreeNode) int {
	ans := 0
	left := make(map[int]int)
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, depth int, position int) {
		if node == nil {
			return
		}
		if _, ok := left[depth]; !ok {
			left[depth] = position
		}
		ans = max(ans, position-left[depth]+1)
		dfs(node.Left, depth+1, position*2)
		dfs(node.Right, depth+1, position*2+1)
	}
	dfs(root, 0, 0)
	return ans
}

// Solution2: BFS
func widthOfBinaryTree1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	type pair struct {
		node    *TreeNode
		positon int
	}
	queue := []pair{{root, 0}}
	for len(queue) > 0 {
		qlen := len(queue)
		var left, right int
		for i := 0; i < qlen; i++ {
			p := queue[0]
			queue = queue[1:]
			node, position := p.node, p.positon
			if i == 0 {
				left = position
			}
			if i == qlen-1 {
				right = position
			}
			if node.Left != nil {
				queue = append(queue, pair{node.Left, position * 2})
			}
			if node.Right != nil {
				queue = append(queue, pair{node.Right, position*2 + 1})
			}
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

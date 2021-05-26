package minimumabsolutedifferenceinbst

import "math"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func minimumAbsDifferenceInBST(root *TreeNode) int {
	min := math.MaxInt64
	prev := -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.left)
		if prev != -1 && (node.val-prev) < min {
			min = node.val - prev
		}
		prev = node.val
		dfs(node.right)
	}
	dfs(root)
	return min
}

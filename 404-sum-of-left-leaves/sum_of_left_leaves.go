package sumofleftleaves

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) (res int) {
	var preOrder func(*TreeNode, bool)
	preOrder = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		if isLeft && node.left == nil && node.right == nil {
			res += node.val
		}
		preOrder(node.left, true)
		preOrder(node.right, false)
	}
	preOrder(root, false)
	return
}

// func isLeafNode(node *TreeNode) bool {
// 	return node.left == nil && node.right == nil
// }

// func dfs(node *TreeNode) (ans int) {
// 	if node.left != nil {
// 		if isLeafNode(node.left) {
// 			ans += node.left.val
// 		} else {
// 			ans += dfs(node.left)
// 		}
// 	}
// 	if node.right != nil && !isLeafNode(node.right) {
// 		ans += dfs(node.right)
// 	}
// 	return
// }

// func sumOfLeftLeaves(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return dfs(root)
// }

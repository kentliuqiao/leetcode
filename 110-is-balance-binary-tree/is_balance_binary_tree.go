package isbalancebinarytree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func isBalanceFromTop(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.left)-height(root.right)) <= 1 &&
		isBalanceFromTop(root.left) &&
		isBalanceFromTop(root.right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.left), height(node.right)) + 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func isBalanceFromBottom(root *TreeNode) bool {
	return heightV2(root) >= 0
}

func heightV2(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := heightV2(node.left)
	rightHeight := heightV2(node.right)
	if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight) + 1
}

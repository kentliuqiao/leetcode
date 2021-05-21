package isbalancetree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(height(root.left)-height(root.right)) <= 1 &&
		isBalanced(root.left) &&
		isBalanced(root.right)
}

func isBalancedFroBottom(root *TreeNode) bool {
	return heightFromBottom(root) >= 0
}

func heightFromBottom(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := heightFromBottom(root.left)
	right := heightFromBottom(root.right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.left), height(root.right)) + 1
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

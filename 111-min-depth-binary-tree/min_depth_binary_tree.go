package mindepthbinarytree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	if root.left == nil {
		return minDepth(root.right) + 1
	} else if root.right == nil {
		return minDepth(root.left) + 1
	}
	return min(minDepth(root.right), minDepth(root.left)) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

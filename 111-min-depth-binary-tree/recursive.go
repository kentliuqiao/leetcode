package mindepthbinarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 168 ms	18 MB
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return min(minDepth(root.Right), minDepth(root.Left)) + 1
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

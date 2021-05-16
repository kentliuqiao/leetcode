package maxdepthnnarytree

type TreeNode struct {
	val      int
	children []*TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, child := range root.children {
		depth := maxDepth(child)
		if max < depth {
			max = depth
		}
	}

	return max + 1
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

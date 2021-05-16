package maxdepthbinarytree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func maxDepth(root *TreeNode) (res int) {
	var preOrder func(n *TreeNode, depth int)
	preOrder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		if depth > res {
			res = depth
		}
		preOrder(node.left, depth)
		preOrder(node.right, depth)
	}
	preOrder(root, 0)

	return
}

func maxDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_1(root.left), maxDepth_1(root.right)) + 1
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func maxDepthLoop(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		ans++
		nextLevel := []*TreeNode{}
		for _, node := range q {
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		q = nextLevel
	}

	return ans
}

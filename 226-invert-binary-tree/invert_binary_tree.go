package invertbinarytree

type TreeNode struct {
	val         int
	Left, Right *TreeNode
}

func invertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		for i, node := range q {
			q[i].Left, q[i].Right = q[i].Right, q[i].Left
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		q = nextLevel
	}

	return root
}

func invertBinaryTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertBinaryTreeRecursive(root.Left)
	right := invertBinaryTreeRecursive(root.Right)
	root.Left = right
	root.Right = left
	return root
}

package invertbinarytree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func invertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		for i, node := range q {
			q[i].left, q[i].right = q[i].right, q[i].left
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
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
	left := invertBinaryTreeRecursive(root.left)
	right := invertBinaryTreeRecursive(root.right)
	root.left = right
	root.right = left
	return root
}

package mergetwotrees

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func mergeTwoTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil && root2 != nil {
		return root2
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	root := &TreeNode{val: root1.val + root2.val}
	root.left = mergeTwoTrees(root1.left, root2.left)
	root.right = mergeTwoTrees(root1.right, root2.right)
	return root
}

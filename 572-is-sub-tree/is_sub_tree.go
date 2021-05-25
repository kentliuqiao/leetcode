package issubtree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func isSubTree(root, subRoot *TreeNode) bool {
	var preOrder func(*TreeNode) bool
	preOrder = func(node *TreeNode) bool {
		return isSameTree(node, subRoot) ||
			isSameTree(node.left, subRoot) ||
			isSameTree(node.right, subRoot)
	}
	return preOrder(root)
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q != nil {
		return false
	}
	return p.val == q.val &&
		isSameTree(p.left, q.left) &&
		isSameTree(p.right, q.right)
}

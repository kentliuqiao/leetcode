package searchinabinarysearchtree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func searchBST(root *TreeNode, target int) *TreeNode {
	if root == nil || root.val == target {
		return root
	}
	if root.val < target {
		return searchBST(root.right, target)
	}
	return searchBST(root.left, target)
}

func searchBSTRecursive(root *TreeNode, target int) *TreeNode {
	for root != nil && root.val != target {
		if root.val < target {
			root = root.right
		} else {
			root = root.left
		}
	}
	return root
}

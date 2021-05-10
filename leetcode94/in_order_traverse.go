package leetcode94

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func inOrderTraverseRecursive(root *TreeNode) (res []int) {
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.left)
		res = append(res, node.val)
		inOrder(node.right)
	}

	inOrder(root)

	return
}

func inOrderTraverseLoop(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, root.val)
		root = root.right
	}

	return
}

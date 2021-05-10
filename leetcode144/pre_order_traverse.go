package leetcode144

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func preOrderTraverseRecursive(root *TreeNode) (res []int) {
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.val)
		preOrder(node.left)
		preOrder(node.right)
	}
	preOrder(root)

	return
}

func preOrderTraverseLoop(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	for root != nil || len(stk) > 0 {
		for root != nil {
			res = append(res, root.val)
			stk = append(stk, root)
			root = root.left
		}
		root = stk[len(stk)-1].right
		stk = stk[:len(stk)-1]
	}

	return
}

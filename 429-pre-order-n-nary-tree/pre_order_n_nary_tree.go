package preorder

func preOrderNnaryTreeRecursive(root *TreeNode) (res []int) {
	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.val)
		for _, child := range node.children {
			preOrder(child)
		}
	}
	preOrder(root)
	return
}

func preOrderNnaryTreeLoop(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	stk := []*TreeNode{root}
	for len(stk) > 0 {
		node := stk[len(stk)-1]
		res = append(res, node.val)
		stk = stk[:len(stk)-1]
		for i := len(node.children) - 1; i >= 0; i-- {
			stk = append(stk, node.children[i])
		}

	}

	return
}

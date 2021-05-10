package leetcode145

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func postOrderRecursive(root *TreeNode) (res []int) {
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.left)
		postOrder(node.right)
		res = append(res, node.val)
	}
	postOrder(root)

	return
}

func postOrderLoop(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode // for the record of the previous accessed node
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.right == nil || prev == root.right {
			res = append(res, root.val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.right
		}
	}

	return
}

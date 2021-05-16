package preordernnarytree

type Node struct {
	val      int
	children []*Node
}

func preOrderNnaryTree(root *Node) (res []int) {
	var preOrder func(*Node)
	preOrder = func(n *Node) {
		if n == nil {
			return
		}
		res = append(res, n.val)
		for _, node := range n.children {
			preOrder(node)
		}
	}
	preOrder(root)

	return
}

func preOrderNnaryTreeLoop(root *Node) (res []int) {
	if root == nil {
		return
	}
	stk := []*Node{root}
	for len(stk) > 0 {
		root = stk[len(stk)-1]
		res = append(res, root.val)
		stk = stk[:len(stk)-1]
		for i := len(root.children) - 1; i >= 0; i-- {
			stk = append(stk, root.children[i])
		}
	}
	return
}

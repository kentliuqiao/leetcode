package postordernnarytree

type Node struct {
	val      int
	children []*Node
}

func postOrderNnaryTree(root *Node) (res []int) {
	var postOrder func(*Node)
	postOrder = func(n *Node) {
		if n == nil {
			return
		}
		for _, child := range n.children {
			postOrder(child)
		}
		res = append(res, n.val)
	}
	postOrder(root)

	return
}

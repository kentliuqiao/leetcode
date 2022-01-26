package connectnextrightnode

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNodde(root.Left, root.Right)
	return root
}

func connectTwoNodde(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectTwoNodde(node1.Left, node1.Right)
	connectTwoNodde(node2.Left, node2.Right)
	connectTwoNodde(node1.Right, node2.Left)
}

package nextrightnode

type TreeNode struct {
	val               int
	left, right, next *TreeNode
}

func nextRightNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			if i == len(q)-1 {
				q[i].next = nil
			} else {
				q[i].next = q[i+1]
			}
			if q[i].left != nil {
				nextLevel = append(nextLevel, q[i].left)
			}
			if q[i].right != nil {
				nextLevel = append(nextLevel, q[i].right)
			}
		}
		q = nextLevel
	}
	return root
}

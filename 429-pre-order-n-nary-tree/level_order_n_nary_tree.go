package preorder

type TreeNode struct {
	val      int
	children []*TreeNode
}

func levelOrderNnaryTree(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range q {
			res = append(res, node.val)
			for _, child := range node.children {
				nextLevel = append(nextLevel, child)
			}
		}
		q = nextLevel
	}
	return
}

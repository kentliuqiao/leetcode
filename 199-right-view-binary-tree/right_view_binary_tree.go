package rightviewbinarytree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		res = append(res, q[len(q)-1].val)
		for _, node := range q {
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		q = nextLevel
	}

	return res
}

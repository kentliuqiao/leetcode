package maxinlevel

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func levelMax(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		max := q[0].val
		for _, node := range q {
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
			if node.val > max {
				max = node.val
			}
		}
		q = nextLevel
		res = append(res, max)
	}
	return
}

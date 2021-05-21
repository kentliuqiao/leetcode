package levelaveragebinarytree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func levelAverage(root *TreeNode) (res []float64) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nextLevel := []*TreeNode{}
		sum, cnt := 0.0, 0.0
		for _, node := range q {
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
			sum += float64(node.val)
			cnt++
		}
		res = append(res, sum/cnt)
		q = nextLevel
	}
	return
}

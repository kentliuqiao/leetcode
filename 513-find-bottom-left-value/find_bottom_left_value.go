package findbottomleftvalue

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	for len(q) > 0 {
		hasNextLevel := false
		nextLevel := []*TreeNode{}
		for _, node := range q {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
				hasNextLevel = true
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
				hasNextLevel = true
			}
		}
		if !hasNextLevel {
			break
		}
		q = nextLevel
	}
	return q[0].Val
}

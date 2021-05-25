package issametree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// DFS
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	return p.val == q.val &&
		isSameTree(p.left, q.left) &&
		isSameTree(p.right, q.right)
}

// BFS
func isSameTreeRecursive(p, q *TreeNode) bool {
	x, y := []*TreeNode{p}, []*TreeNode{q}
	for len(x) > 0 {
		m, n := x[0], y[0]
		x, y = x[1:], y[1:]
		if m == nil && n == nil {
			continue
		}
		if m == nil && n != nil {
			return false
		}
		if m != nil && n == nil {
			return false
		}
		if m.val != n.val {
			return false
		}
		x = append(x, m.left, m.right)
		y = append(y, n.left, n.right)
	}
	return true
}

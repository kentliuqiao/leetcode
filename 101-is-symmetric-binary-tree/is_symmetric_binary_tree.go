package issymmetric

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	return left.val == right.val &&
		check(left.left, right.right) &&
		check(left.right, right.left)
}

func isSymmetricLoop(root *TreeNode) bool {
	u, v := root, root
	q := []*TreeNode{}
	q = append(q, []*TreeNode{u, v}...)
	for len(q) > 0 {
		u, v := q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.val != v.val {
			return false
		}
		q = append(q, u.left)
		q = append(q, v.right)

		q = append(q, u.right)
		q = append(q, v.left)
	}

	return true
}

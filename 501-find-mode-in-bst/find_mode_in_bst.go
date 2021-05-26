package findmodeinbst

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func findMode(root *TreeNode) (res []int) {
	base, count, maxCount := 0, 0, 0
	// Note: it must used under the circumancetance of a sorted list
	update := func(x int) {
		if base == x {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}
	// vals := []int{}
	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.left)

		update(node.val)

		inOrder(node.right)
	}
	inOrder(root)

	// for _, v := range vals {
	// 	update(v)
	// }
	return
}

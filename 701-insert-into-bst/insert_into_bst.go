package insertintobst

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	p := root
	for p != nil {
		if p.val < val {
			if p.right == nil {
				p.right = &TreeNode{val: val}
				break
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = &TreeNode{val: val}
				break
			}
			p = p.left
		}
	}

	return root
}

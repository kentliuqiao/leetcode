package leetcode102_right_view_of_tree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func rightViewOfBinaryTree(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		temp := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			if j == len(q)-1 {
				ret = append(ret, node.val)
			}
			if node.left != nil {
				temp = append(temp, node.left)
			}
			if node.right != nil {
				temp = append(temp, node.right)
			}
		}
		q = temp
	}

	return ret
}

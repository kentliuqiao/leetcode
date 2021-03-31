package leetcode102

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{}) // init a row in array
		temp := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.val)
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

package leetcode107

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func levelOrderBottom(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	res = [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		tmp := []int{}
		for _, node := range queue {
			tmp = append(tmp, node.val)
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		res = append(res, tmp)
		queue = nextLevel
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return
}

func levelOrder(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range queue {
			res = append(res, node.val)
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		queue = nextLevel
	}
	return
}

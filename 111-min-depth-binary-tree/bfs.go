package mindepthbinarytree

// 188 ms	18.8 MB
func bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
		q = q[size:]
	}
	return depth
}

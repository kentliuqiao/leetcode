package maximumdepthofbinarytree

// traverse
// 遍历一遍二叉树
// 4 ms	4.2 MB
func traverse(root *TreeNode) int {
	depth, res := 0, 0
	var traverseFunc = func(*TreeNode) {}
	traverseFunc = func(root *TreeNode) {
		if root == nil {
			res = max(res, depth)
			return
		}
		depth++
		traverseFunc(root.Left)
		traverseFunc(root.Right)
		depth--
	}
	traverseFunc(root)
	return res
}

package diameterofbinarytree

// 执行用时：4 ms , 在所有 Go 提交中击败了 91.98% 的用户
// 内存消耗：4.4 MB, 在所有 Go 提交中击败了 58.53% 的用户
// 遇到子树问题，首先想到的是给函数设置返回值，然后在后序位置做文章。
func divideConquer(root *TreeNode) int {
	res := 0
	var maxDepth func(*TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		res = max(res, left+right)
		return 1 + max(left, right)
	}
	maxDepth(root)
	return res
}

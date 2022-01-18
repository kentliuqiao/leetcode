package maximumdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 	4 ms	4.2 MB
// divide and conquer
// 分解问题：一颗二叉树的最大高度可以由左右的子树的最大高度得到
func recursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := recursive(root.Left)
	rightMax := recursive(root.Right)
	// 后序位置
	// 只有在这个位置，才能拿到左右子树的最大高度！
	return max(leftMax, rightMax) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

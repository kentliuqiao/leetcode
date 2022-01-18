package diameterofbinarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 执行用时：20 ms, 在所有 Go 提交中击败了 7.57% 的用户
// 内存消耗：4.4 MB , 在所有 Go 提交中击败了 10.40% 的用户
func recursive(root *TreeNode) int {
	res := 0
	var recursiveFunc func(*TreeNode)
	recursiveFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		leftMaxDepth := maxDepth(root.Left)
		rightMaxDepth := maxDepth(root.Right)
		diameter := leftMaxDepth + rightMaxDepth
		res = max(res, diameter)

		recursiveFunc(root.Left)
		recursiveFunc(root.Right)
	}
	recursiveFunc(root)
	return res
}

func maxDepth(root *TreeNode) int {
	res := 0
	depth := 0
	var depthFunc func(*TreeNode)
	depthFunc = func(root *TreeNode) {
		if root == nil {
			res = max(res, depth)
			return
		}
		depth++
		depthFunc(root.Left)
		depthFunc(root.Right)
		depth--
	}
	depthFunc(root)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

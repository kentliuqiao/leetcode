package maximumbinarytree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}
	largest := math.MinInt64
	index := -1
	for i := lo; i <= hi; i++ {
		if largest < nums[i] {
			index = i
			largest = nums[i]
		}
	}
	if largest == math.MinInt64 || index == -1 {
		return nil
	}

	node := &TreeNode{Val: largest}
	left := build(nums, lo, index-1)
	right := build(nums, index+1, hi)
	node.Left = left
	node.Right = right
	return node
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

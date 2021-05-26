package isbinarysearchtree

import (
	"math"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func isBinarySearchTree(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.val <= lower || root.val >= upper {
		return false
	}
	return helper(root.left, lower, root.val) && helper(root.right, root.val, upper)
}

// Accordong to BST's definition, when we traverse it by inOrder, we can get a list of values with ascending sort order
func isBinarySearchTreeInOrder(root *TreeNode) bool {
	stk := []*TreeNode{}
	prev := math.MinInt64
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.val <= prev {
			return false
		}
		prev = root.val
		root = root.right
	}
	return true
}

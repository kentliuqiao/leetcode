package isbalancetree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &TreeNode{
	val: 12,
	left: &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 83,
			right: &TreeNode{
				val: 5,
			},
		},
	},
	right: &TreeNode{
		val: 4,
		right: &TreeNode{
			val: 0,
			left: &TreeNode{
				val: 9,
			},
		},
	},
}

var cases = []struct {
	tree *TreeNode
	want bool
}{
	{
		tree: tree,
		want: false,
	},
	{
		tree: nil,
		want: true,
	},
	{
		tree: &TreeNode{val: 1},
		want: true,
	},
	{
		tree: &TreeNode{val: 1, left: &TreeNode{}, right: &TreeNode{}},
		want: true,
	},
}

func TestIsBalancedTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isBalanced(c.tree))
	}
}

func TestIsBalancedTreeFromBottom(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isBalancedFroBottom(c.tree))
	}
}

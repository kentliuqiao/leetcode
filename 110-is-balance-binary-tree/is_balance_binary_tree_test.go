package isbalancebinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tree = &TreeNode{
		val: 12,
		left: &TreeNode{
			val: 3,
			left: &TreeNode{
				val: 83,
			},
			right: &TreeNode{
				val: 5,
			},
		},
		right: &TreeNode{
			val: 4,
			right: &TreeNode{
				val: 0,
			},
		},
	}
	treeUnbalance = &TreeNode{
		val: 1,
		right: &TreeNode{
			val: 2,
			right: &TreeNode{
				val: 3,
				right: &TreeNode{
					val: 4,
				},
			},
		},
	}
)

var cases = []struct {
	tree *TreeNode
	want bool
}{
	{
		tree: tree,
		want: true,
	},
	{
		tree: treeUnbalance,
		want: false,
	},
	{
		tree: &TreeNode{val: -1},
		want: true,
	},
	{
		tree: nil,
		want: true,
	},
}

func TestIsBalanceFromTop(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isBalanceFromTop(c.tree))
	}
}

func TestIsBalanceFromBottom(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isBalanceFromBottom(c.tree))
	}
}

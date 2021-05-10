package leetcode144

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree1 = &TreeNode{
	val: 8,
	left: &TreeNode{
		val: 2,
		left: &TreeNode{
			val: 90,
			right: &TreeNode{
				val: 5,
			},
		},
	},
	right: &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 10,
			left: &TreeNode{
				val: 1,
			},
			right: &TreeNode{
				val: 23,
				left: &TreeNode{
					val: -12,
				},
			},
		},
	},
}

var cases = []struct {
	tree *TreeNode
	want []int
}{
	{
		tree: nil,
		want: []int(nil),
	},
	{
		tree: tree1,
		want: []int{8, 2, 90, 5, 3, 10, 1, 23, -12},
	},
}

func TestPreOrderRecursive(t *testing.T) {
	for _, c := range cases {
		got := preOrderTraverseRecursive(c.tree)
		assert.Equal(t, c.want, got)
	}
}

func TestPreOrderLoop(t *testing.T) {
	for _, c := range cases {
		got := preOrderTraverseLoop(c.tree)
		assert.Equal(t, c.want, got)
	}
}

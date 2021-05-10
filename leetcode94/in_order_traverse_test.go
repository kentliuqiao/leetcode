package leetcode94

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
		want: []int{90, 5, 2, 8, 1, 10, -12, 23, 3},
	},
}

func TestInOrderTraverseRecursive(t *testing.T) {
	for _, c := range cases {
		got := inOrderTraverseRecursive(c.tree)
		assert.Equal(t, c.want, got)
	}
}

func TestInOrderTraverseLoop(t *testing.T) {
	for _, c := range cases {
		got := inOrderTraverseLoop(c.tree)
		assert.Equal(t, c.want, got)
	}
}

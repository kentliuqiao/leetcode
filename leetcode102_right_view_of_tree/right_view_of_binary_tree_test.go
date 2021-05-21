package leetcode102_right_view_of_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree1 *TreeNode = &TreeNode{
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
		},
	},
}

var cases = []struct {
	tree  *TreeNode
	want1 []int
	want2 []int
}{
	{
		tree:  tree1,
		want1: []int{12, 4, 0, 5},
		want2: []int{12, 3, 4, 83, 0, 5},
	},
	{
		tree:  nil,
		want1: []int(nil),
		want2: []int(nil),
	},
}

func TestRightViewOfBinaryTree(t *testing.T) {
	for _, c := range cases {
		got := rightViewOfBinaryTree(c.tree)
		assert.Equal(t, c.want1, got)
	}
}

func TestLevelOrder(t *testing.T) {
	for _, c := range cases {
		got := levelOrder(c.tree)
		assert.Equal(t, c.want2, got)
	}
}

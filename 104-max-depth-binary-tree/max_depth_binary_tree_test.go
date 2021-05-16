package maxdepthbinarytree

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
		},
	},
}

var cases = []struct {
	tree *TreeNode
	want int
}{
	{
		tree: tree,
		want: 4,
	},
	{
		tree: &TreeNode{val: 1},
		want: 1,
	},
	{
		tree: nil,
		want: 0,
	},
}

func TestMaxDepthBinaryTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, maxDepth(c.tree))
	}
}

func TestMaxDepthBinaryTree_1(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, maxDepth_1(c.tree))
	}
}

func TestMaxDepthBinaryTreeLoop(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, maxDepthLoop(c.tree))
	}
}

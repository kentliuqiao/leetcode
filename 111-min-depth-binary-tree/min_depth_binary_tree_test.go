package mindepthbinarytree

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
			val: 1,
		},
	},
}

var cases = []struct {
	tree *TreeNode
	want int
}{
	{
		tree: tree,
		want: 3,
	},
	{
		tree: &TreeNode{val: 1, left: &TreeNode{val: 2}},
		want: 2,
	},
	{
		tree: nil,
		want: 0,
	},
}

func TestMinDepth(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, minDepth(c.tree))
	}
}

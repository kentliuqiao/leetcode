package sumofleftleaves

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
	want int
}{
	{
		tree: tree,
		want: 9,
	},
	{
		tree: &TreeNode{
			val:   1,
			left:  &TreeNode{val: 2},
			right: &TreeNode{val: 3},
		},
		want: 2,
	},
	{
		tree: &TreeNode{
			val:   1,
			right: &TreeNode{val: 2},
		},
		want: 0,
	},
	{
		tree: &TreeNode{val: 21},
		want: 0,
	},
	{
		tree: nil,
		want: 0,
	},
}

func TestSumOfLeftLeaves(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, sumOfLeftLeaves(c.tree))
	}
}

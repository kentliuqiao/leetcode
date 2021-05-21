package preorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &TreeNode{
	val: 1,
	children: []*TreeNode{
		{
			val: 32,
			children: []*TreeNode{
				{
					val: -12,
					children: []*TreeNode{
						{
							val: 333,
						},
					},
				},
			},
		},
		{
			val: 11,
			children: []*TreeNode{
				{
					val: 19,
				},
				{
					val: 54,
				},
				{
					val: 1222,
				},
			},
		},
		{
			val: 123,
			children: []*TreeNode{
				{
					val: 34,
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
		tree: tree,
		want: []int{1, 32, 11, 123, -12, 19, 54, 1222, 34, 333},
	},
	{
		tree: nil,
		want: []int(nil),
	},
	{
		tree: &TreeNode{},
		want: []int{0},
	},
}

func TestLevelOrderNnaryTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, levelOrderNnaryTree(c.tree))
	}
}

package maxdepthnnarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &TreeNode{
	val: 12,
	children: []*TreeNode{
		{
			val: 9,
			children: []*TreeNode{
				{
					val: 1,
					children: []*TreeNode{
						{val: 6},
						{val: -1},
					},
				},
			},
		},
		{
			val: 4,
			children: []*TreeNode{
				{val: 11},
				{val: 3},
			},
		},
		{
			val: 7,
		},
		{
			val: 6,
		},
	},
}

var cases = []struct {
	tree  *TreeNode
	depth int
}{
	{
		tree:  tree,
		depth: 4,
	},
	{
		tree:  &TreeNode{val: 1},
		depth: 1,
	},
	{tree: nil, depth: 0},
}

func TestMaxDepthNnaryTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.depth, maxDepth(c.tree))
	}
}

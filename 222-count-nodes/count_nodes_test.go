package countnodes

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
		want: 7,
	},
	{
		tree: nil,
		want: 0,
	},
	{
		tree: &TreeNode{val: 1},
		want: 1,
	},
}

func TestCountNodes(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, countNodes(c.tree))
	}
}

func TestCountNodesPreOrder(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, countNodesPreOrderRecursive(c.tree))
	}
}

func TestCountNodes_Optimization(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, countNodesOptimization(c.tree))
	}
}

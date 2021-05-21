package levelaveragebinarytree

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
	want []float64
}{
	{
		tree: tree,
		want: []float64{12, 3.5, 41.5, 7},
	},
	{
		tree: nil,
		want: []float64(nil),
	},
	{
		tree: &TreeNode{val: 1},
		want: []float64{1},
	},
}

func TestLevelAverage(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, levelAverage(c.tree))
	}
}

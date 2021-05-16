package maxinlevel

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
	want []int
}{
	{
		tree: tree,
		want: []int{12, 4, 83, 5},
	},
	{
		tree: &TreeNode{val: -1},
		want: []int{-1},
	},
	{
		tree: nil,
		want: []int(nil),
	},
}

func TestMaxInLevel(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, levelMax(c.tree))
	}
}

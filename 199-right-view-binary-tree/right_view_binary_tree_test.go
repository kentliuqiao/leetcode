package rightviewbinarytree

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
	want []int
}{
	{
		tree: tree,
		want: []int{12, 4, 0, 9},
	},
	{
		tree: nil,
		want: []int(nil),
	},
	{
		tree: &TreeNode{val: 1},
		want: []int{1},
	},
}

func TestRightSideView(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, rightSideView(c.tree))
	}
}

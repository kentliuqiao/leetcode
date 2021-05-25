package findbottomleftvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &TreeNode{
	Val: 12,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 83,
			Right: &TreeNode{
				Val: 5,
			},
		},
	},
	Right: &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 9,
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
		want: 5,
	},
	{
		tree: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		want: 2,
	},
	{
		tree: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		want: 2,
	},
	{
		tree: &TreeNode{Val: 21},
		want: 21,
	},
}

func TestFindBottomLeftValue(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, findBottomLeftValue(c.tree))
	}
}

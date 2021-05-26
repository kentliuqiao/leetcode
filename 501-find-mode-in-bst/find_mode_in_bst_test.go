package findmodeinbst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	tree *TreeNode
	want []int
}{
	{
		tree: &TreeNode{val: 4, left: &TreeNode{val: 4, left: &TreeNode{val: 1}, right: &TreeNode{val: 4}}, right: &TreeNode{val: 7}},
		want: []int{4},
	},
	{
		tree: &TreeNode{val: 8, left: &TreeNode{val: 6, left: &TreeNode{val: 6}, right: &TreeNode{val: 8}}},
		want: []int{6, 8},
	},
	{
		tree: &TreeNode{val: 1},
		want: []int{1},
	},
	{
		tree: nil,
		want: []int(nil),
	},
}

func TestFindModeInBST(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, findMode(c.tree))
	}
}

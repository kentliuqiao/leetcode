package mergetwotrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	root1, root2 *TreeNode
	want         *TreeNode
}{
	{
		root1: &TreeNode{val: 1, left: &TreeNode{val: 3, left: &TreeNode{val: 5}}, right: &TreeNode{val: 2}},
		root2: &TreeNode{val: 2, left: &TreeNode{val: 1, right: &TreeNode{val: 4}}, right: &TreeNode{val: 3, right: &TreeNode{val: 7}}},
		want:  &TreeNode{val: 3, left: &TreeNode{4, &TreeNode{val: 5}, &TreeNode{val: 4}}, right: &TreeNode{val: 5, right: &TreeNode{val: 7}}},
	},
	{
		root1: &TreeNode{val: 2, right: &TreeNode{val: 4, right: &TreeNode{val: 6}}},
		root2: &TreeNode{val: 3, left: &TreeNode{val: 5, left: &TreeNode{val: 7}}},
		want:  &TreeNode{val: 5, left: &TreeNode{val: 5, left: &TreeNode{val: 7}}, right: &TreeNode{val: 4, right: &TreeNode{val: 6}}},
	},
	{
		root1: &TreeNode{val: 2, right: &TreeNode{val: 1}},
		root2: nil,
		want:  &TreeNode{val: 2, right: &TreeNode{val: 1}},
	},
	{
		root1: nil,
		root2: &TreeNode{val: 1, left: &TreeNode{val: 0}},
		want:  &TreeNode{val: 1, left: &TreeNode{val: 0}},
	},
	{
		root1: nil,
		root2: nil,
		want:  nil,
	},
}

func TestMergeTwoTrees(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, mergeTwoTrees(c.root1, c.root2))
	}
}

package searchinabinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	tree   *TreeNode
	target int
	want   *TreeNode
}{
	{
		tree:   &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7}},
		target: 2,
		want:   &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}},
	},
	{
		tree:   &TreeNode{val: 3, left: &TreeNode{val: 2}, right: &TreeNode{val: 7}},
		target: 7,
		want:   &TreeNode{val: 7},
	},
	{
		tree:   &TreeNode{val: 3, left: &TreeNode{val: 2}, right: &TreeNode{val: 7}},
		target: 3,
		want:   &TreeNode{val: 3, left: &TreeNode{val: 2}, right: &TreeNode{val: 7}},
	},
	{
		tree:   &TreeNode{val: 3, left: &TreeNode{val: 2}, right: &TreeNode{val: 7}},
		target: -1,
		want:   nil,
	},
	{
		tree:   nil,
		target: 0,
		want:   nil,
	},
}

func TestSearchBST(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, searchBST(c.tree, c.target))
	}
}

func TestSearchBSTRecursive(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, searchBSTRecursive(c.tree, c.target))
	}
}

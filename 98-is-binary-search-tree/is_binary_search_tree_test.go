package isbinarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	tree *TreeNode
	want bool
}{
	{
		tree: &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7}},
		want: true,
	},
	{
		tree: &TreeNode{val: 4, left: &TreeNode{val: 6, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7}},
		want: false,
	},
	{
		tree: &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 5}}, right: &TreeNode{val: 7}},
		want: false,
	},
	{
		tree: &TreeNode{val: 2},
		want: true,
	},
	{
		tree: nil,
		want: true,
	},
	{
		tree: &TreeNode{val: 3, right: &TreeNode{val: 9}},
		want: true,
	},
}

func TestIsBinarySearchTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isBinarySearchTree(c.tree))
	}
}

func TestIsBinarySearchTreeInOrder(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isBinarySearchTreeInOrder(c.tree))
	}
}

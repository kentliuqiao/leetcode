package buildtreefrompreinorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	inOrder, preOrder []int
	want              *TreeNode
}{
	{
		inOrder:  []int{9, 3, 15, 20, 7},
		preOrder: []int{3, 9, 20, 15, 7},
		want: &TreeNode{
			val:   3,
			left:  &TreeNode{val: 9},
			right: &TreeNode{val: 20, left: &TreeNode{val: 15}, right: &TreeNode{val: 7}},
		},
	},
	{
		inOrder:  []int{1},
		preOrder: []int{1},
		want:     &TreeNode{val: 1},
	},
	{
		inOrder:  []int{2, 3, 4},
		preOrder: []int{4, 3, 2},
		want:     &TreeNode{val: 4, left: &TreeNode{val: 3, left: &TreeNode{val: 2}}},
	},
	{
		inOrder:  []int{},
		preOrder: []int{},
		want:     nil,
	},
}

func TestBuildTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, buildTree(c.preOrder, c.inOrder))
	}
}

func TestBuildTreeV2(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, buildTreeV2(c.preOrder, c.inOrder))
	}
}

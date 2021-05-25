package buildtreefrominpostorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	inOrder, postOrder []int
	want               *TreeNode
}{
	{
		inOrder:   []int{9, 3, 15, 20, 7},
		postOrder: []int{9, 15, 7, 20, 3},
		want: &TreeNode{
			val:   3,
			left:  &TreeNode{val: 9},
			right: &TreeNode{val: 20, left: &TreeNode{val: 15}, right: &TreeNode{val: 7}},
		},
	},
	{
		inOrder:   []int{1},
		postOrder: []int{1},
		want:      &TreeNode{val: 1},
	},
	{
		inOrder:   []int{4, 3, 2},
		postOrder: []int{4, 3, 2},
		want:      &TreeNode{val: 2, left: &TreeNode{val: 3, left: &TreeNode{val: 4}}},
	},
	{
		inOrder:   []int{},
		postOrder: []int{},
		want:      nil,
	},
}

func TestBuildTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, buildTree(c.inOrder, c.postOrder))
	}
}

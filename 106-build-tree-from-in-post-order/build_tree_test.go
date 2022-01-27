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
			Val:   3,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		},
	},
	{
		inOrder:   []int{1},
		postOrder: []int{1},
		want:      &TreeNode{Val: 1},
	},
	{
		inOrder:   []int{4, 3, 2},
		postOrder: []int{4, 3, 2},
		want:      &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
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

func TestBuildTreeV2(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, buildTreeV2(c.inOrder, c.postOrder))
	}
}

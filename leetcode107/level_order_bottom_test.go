package leetcode107

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
	tree       *TreeNode
	want       []int
	wantBottom [][]int
}{
	{
		tree:       tree,
		want:       []int{12, 3, 4, 83, 0, 5, 9},
		wantBottom: [][]int{{5, 9}, {83, 0}, {3, 4}, {12}},
	},
	{
		tree:       nil,
		want:       []int(nil),
		wantBottom: [][]int(nil),
	},
}

func TestLevelOrder(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, levelOrder(c.tree))
	}
}

func TestLevelOrderBottom(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.wantBottom, levelOrderBottom(c.tree))
	}
}

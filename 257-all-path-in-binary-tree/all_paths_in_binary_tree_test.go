package allpathsinbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tree = &TreeNode{
		val: 12,
		left: &TreeNode{
			val: 3,
			left: &TreeNode{
				val: 83,
			},
			right: &TreeNode{
				val: 5,
			},
		},
		right: &TreeNode{
			val: 4,
			right: &TreeNode{
				val: 0,
			},
		},
	}
)

var cases = []struct {
	tree *TreeNode
	want []string
}{
	{
		tree: tree,
		want: []string{"12->3->83", "12->3->5", "12->4->0"},
	},
	{
		tree: &TreeNode{val: 1, right: &TreeNode{val: 2}},
		want: []string{"1->2"},
	},
	{
		tree: &TreeNode{val: -1},
		want: []string{"-1"},
	},
	{
		tree: nil,
		want: []string{},
	},
}

func TestAllPaths(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, binaryTreePaths(c.tree))
	}
}

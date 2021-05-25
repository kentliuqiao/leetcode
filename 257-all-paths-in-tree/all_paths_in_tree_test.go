package allpathsintree

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
	tree *TreeNode
	want []string
}{
	{
		tree: tree,
		want: []string{"12->3->83->5", "12->4->0->9"},
	},
	{
		tree: &TreeNode{
			val:   1,
			left:  &TreeNode{val: 2},
			right: &TreeNode{val: 3},
		},
		want: []string{"1->2", "1->3"},
	},
	{
		tree: &TreeNode{
			val:   1,
			right: &TreeNode{val: 2},
		},
		want: []string{"1->2"},
	},
	{
		tree: &TreeNode{val: 1},
		want: []string{"1"},
	},
	{
		tree: nil,
		want: []string{},
	},
}

func TestAllPaths(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, allPaths(c.tree))
	}
}

func TestAllPathsRecursive(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, allPathsRecursive(c.tree))
	}
}

package lowestcommonancestor

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
		},
	},
}

var cases = []struct {
	tree, p, q *TreeNode
	want       *TreeNode
	name       string
}{
	{
		tree: tree, p: &TreeNode{val: 3}, q: &TreeNode{val: 4},
		want: &TreeNode{val: 12},
		name: "case1",
	},
	{
		tree: tree, p: &TreeNode{val: 5}, q: &TreeNode{val: 3},
		want: &TreeNode{val: 3},
		name: "case2",
	},
	{
		tree: tree, p: &TreeNode{val: 4}, q: &TreeNode{val: 0},
		want: &TreeNode{val: 4},
		name: "case3",
	},
	{
		tree: nil, p: nil, q: nil,
		want: nil,
		name: "case4",
	},
	{
		tree: tree, p: &TreeNode{val: 4}, q: &TreeNode{val: 4},
		want: &TreeNode{val: 4},
		name: "case5",
	},
}

func TestLowestCommonAncestor(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, lowestCommonAncestor(c.tree, c.p, c.q))
		})
	}
}

func TestLowestCommonAncestorMap(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.want, lowestCommonAncestorMap(c.tree, c.p, c.q))
		})
	}
}

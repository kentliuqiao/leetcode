package mindepthbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &TreeNode{
	Val: 12,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 83,
			Right: &TreeNode{
				Val: 5,
			},
		},
	},
	Right: &TreeNode{
		Val: 4,
		Right: &TreeNode{
			Val: 1,
		},
	},
}

var cases = []struct {
	tree *TreeNode
	want int
}{
	{
		tree: tree,
		want: 3,
	},
	{
		tree: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		want: 2,
	},
	{
		tree: nil,
		want: 0,
	},
}

func TestMinDepth(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, minDepth(c.tree))
		assert.Equal(t, c.want, bfs(c.tree))
	}
}

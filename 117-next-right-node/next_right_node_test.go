package nextrightnode

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

var wantTree = &TreeNode{
	val: 12,
	left: &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 83,
			right: &TreeNode{
				val: 5,
			},
			next: &TreeNode{
				val: 0,
			},
		},
		next: &TreeNode{
			val: 4,
		},
	},
	right: &TreeNode{
		val: 4,
		right: &TreeNode{
			val: 0,
		},
	},
	next: nil,
}

var cases = []struct {
	tree *TreeNode
	want *TreeNode
}{
	{
		tree: tree,
		want: wantTree,
	},
	{
		tree: &TreeNode{val: -1},
		want: &TreeNode{val: -1},
	},
	{
		tree: nil,
		want: nil,
	},
}

func TestNextRightNode(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, nextRightNode(c.tree))
	}
}

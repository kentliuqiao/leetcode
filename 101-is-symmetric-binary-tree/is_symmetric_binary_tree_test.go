package issymmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	symmetricTree = &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,

			right: &TreeNode{
				val: 7,
			},
		},
		right: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 7,
			},
		},
	}
	nonSymmetricTree = &TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 4,
			},
			right: &TreeNode{
				val: 3,
			},
		},
		right: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 7,
			},
			right: &TreeNode{
				val: 4,
			},
		},
	}

	cases = []struct {
		tree *TreeNode
		want bool
	}{
		{
			tree: symmetricTree,
			want: true,
		},
		{
			tree: nonSymmetricTree,
			want: false,
		},
		{
			tree: &TreeNode{val: 3},
			want: true,
		},
		{
			tree: &TreeNode{val: 1, left: &TreeNode{val: 2}},
			want: false,
		},
		{
			tree: &TreeNode{},
			want: true,
		},
		{
			tree: nil,
			want: true,
		},
	}
)

func TestIsSymmetric(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isSymmetric(c.tree))
	}
}

func TestIsSymmetricLoop(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isSymmetricLoop(c.tree))
	}
}

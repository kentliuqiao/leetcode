package issametree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	p, q *TreeNode
	want bool
}{
	{
		p: &TreeNode{
			val: 1,
			left: &TreeNode{
				val: 2,
				right: &TreeNode{
					val: 3,
				},
			},
			right: &TreeNode{
				val: 4,
				right: &TreeNode{
					val: 5,
				},
			},
		},
		q: &TreeNode{
			val: 1,
			left: &TreeNode{
				val: 2,
				right: &TreeNode{
					val: 3,
				},
			},
			right: &TreeNode{
				val: 4,
				right: &TreeNode{
					val: 5,
				},
			},
		},
		want: true,
	},
	{
		p: &TreeNode{
			val:  1,
			left: &TreeNode{val: 2},
		},
		q: &TreeNode{
			val:  1,
			left: &TreeNode{val: 3},
		},
		want: false,
	},
	{
		p: &TreeNode{
			val:  1,
			left: &TreeNode{val: 2},
		},
		q: &TreeNode{
			val:   1,
			right: &TreeNode{val: 2},
		},
		want: false,
	},
	{
		p:    nil,
		q:    nil,
		want: true,
	},
}

func TestIsSameTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isSameTree(c.p, c.q))
	}
}

func TestIsSameTreeRecursive(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, isSameTreeRecursive(c.p, c.q))
	}
}

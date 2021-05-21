package leetcode145

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ interface{} = map[string]string(nil) // notice: this map variable cannot be used to set key value pair, it must be initialized with make function specifically.
	_ interface{} = []int(nil)             // this slice of int variable can be used to insert with append function but cannot used to insert with index directly.
	_ interface{} = (*TreeNode)(nil)
	_ interface{} = (*string)(nil)
	_ chan int    = (chan int)(nil)
)

var tree1 = &TreeNode{
	val: 8,
	left: &TreeNode{
		val: 2,
		left: &TreeNode{
			val: 90,
			right: &TreeNode{
				val: 5,
			},
		},
	},
	right: &TreeNode{
		val: 3,
		left: &TreeNode{
			val: 10,
			left: &TreeNode{
				val: 1,
			},
			right: &TreeNode{
				val: 23,
				left: &TreeNode{
					val: -12,
				},
			},
		},
	},
}

var cases = []struct {
	tree *TreeNode
	want []int
}{
	{
		tree: nil,
		want: []int(nil),
	},
	{
		tree: tree1,
		want: []int{5, 90, 2, 1, -12, 23, 10, 3, 8},
	},
}

func TestPostOrderRecursive(t *testing.T) {
	for _, c := range cases {
		got := postOrderRecursive(c.tree)
		assert.Equal(t, c.want, got)
	}
}

func TestPostOrderLoop(t *testing.T) {
	for _, c := range cases {
		got := postOrderLoop(c.tree)
		assert.Equal(t, c.want, got)
	}
}

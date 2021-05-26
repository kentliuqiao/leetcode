package buildmaximumtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums []int
	want *TreeNode
}{
	{
		nums: []int{3, 2, 1, 6, 0, 5},
		want: &TreeNode{
			val:   6,
			left:  &TreeNode{val: 3, right: &TreeNode{val: 2, right: &TreeNode{val: 1}}},
			right: &TreeNode{val: 5, left: &TreeNode{val: 0}},
		},
	},
	{
		nums: []int{3, 2, 1},
		want: &TreeNode{val: 3, right: &TreeNode{val: 2, right: &TreeNode{val: 1}}},
	},
	{
		nums: []int{3},
		want: &TreeNode{val: 3},
	},
	{
		nums: []int{},
		want: nil,
	},
	{
		nums: nil,
		want: nil,
	},
}

func TestBuildMaximumTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, builMaximumTree(c.nums))
	}
}

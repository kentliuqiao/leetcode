package preorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var casesPreOrder = []struct {
	tree *TreeNode
	want []int
}{
	{
		tree: tree,
		want: []int{1, 32, -12, 333, 11, 19, 54, 1222, 123, 34},
	},
	{
		tree: &TreeNode{val: 1},
		want: []int{1},
	},
	{
		tree: nil,
		want: []int(nil),
	},
}

func TestPreOrderNnaryTree(t *testing.T) {
	for _, c := range casesPreOrder {
		assert.Equal(t, c.want, preOrderNnaryTreeRecursive(c.tree))
	}
}

func TestPreOrderNnaryTreeLoop(t *testing.T) {
	for _, c := range casesPreOrder {
		assert.Equal(t, c.want, preOrderNnaryTreeLoop(c.tree))
	}
}

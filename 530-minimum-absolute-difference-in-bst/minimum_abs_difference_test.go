package minimumabsolutedifferenceinbst

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	tree *TreeNode
	want int
}{
	{
		tree: &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7}},
		want: 1,
	},
	{
		tree: &TreeNode{val: 9, left: &TreeNode{val: 2, right: &TreeNode{val: 4}}},
		want: 2,
	},
	{
		tree: nil,
		want: math.MaxInt64,
	},
}

func TestMinimumAbsDifference(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, minimumAbsDifferenceInBST(c.tree))
	}
}

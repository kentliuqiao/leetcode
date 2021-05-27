package insertintobst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	root *TreeNode
	val  int
}

var tree = &TreeNode{
	val: 4,
	left: &TreeNode{
		val:   2,
		left:  &TreeNode{val: 1},
		right: &TreeNode{val: 3},
	},
	right: &TreeNode{
		val: 7,
	},
}

var tests = []struct {
	name string
	args args
	want *TreeNode
}{
	{
		name: "cases1",
		args: args{
			root: tree,
			val:  8,
		},
		want: &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7, right: &TreeNode{val: 8}}},
	},
	{
		name: "case2",
		args: args{
			root: &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7}},
			val:  5,
		},
		want: &TreeNode{val: 4, left: &TreeNode{val: 2, left: &TreeNode{val: 1}, right: &TreeNode{val: 3}}, right: &TreeNode{val: 7, left: &TreeNode{val: 5}}},
	},
	{
		name: "case3",
		args: args{
			root: nil,
			val:  1,
		},
		want: &TreeNode{val: 1},
	},
}

func Test_insertIntoBST(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, insertIntoBST(tt.args.root, tt.args.val))
		})
	}
}

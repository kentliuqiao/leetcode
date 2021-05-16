package preordernnarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tree = &Node{
	val: 12,
	children: []*Node{
		{
			val: 9,
			children: []*Node{
				{
					val: 1,
					children: []*Node{
						{val: 6},
						{val: -1},
					},
				},
			},
		},
		{
			val: 4,
			children: []*Node{
				{val: 11},
				{val: 3},
			},
		},
		{
			val: 7,
		},
		{
			val: 6,
		},
	},
}

var cases = []struct {
	tree *Node
	want []int
}{
	{
		tree: tree,
		want: []int{12, 9, 1, 6, -1, 4, 11, 3, 7, 6},
	},
	{
		tree: &Node{val: 1},
		want: []int{1},
	},
	{
		tree: nil,
		want: []int(nil),
	},
}

func TestPreOrderNnaryTree(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, preOrderNnaryTree(c.tree))
	}
}

func TestPreOrderNnaryTreeLoop(t *testing.T) {
	for _, c := range cases {
		assert.Equal(t, c.want, preOrderNnaryTreeLoop(c.tree))
	}
}

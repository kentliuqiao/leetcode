package swappairslinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	head *ListNode
}

var tests = []struct {
	name string
	args args
	want *ListNode
}{
	{
		name: "case1",
		args: args{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
		want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
	},
	{
		name: "case2",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
		},
		want: &ListNode{Val: 3, Next: &ListNode{Val: 2}},
	},
	{
		name: "case3",
		args: args{
			head: nil,
		},
		want: nil,
	},
}

func Test_swapPairs(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, swapPairs(tt.args.head))
		})
	}
}

func Test_swapPairsRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, swapPairsRecursive(tt.args.head))
		})
	}
}

package removenthfromendlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	head *ListNode
	n    int
}

var tests = []struct {
	name string
	args args
	want *ListNode
}{
	{
		name: "case1",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}},
			n:    2,
		},
		want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}},
	},
	{
		name: "case2",
		args: args{
			head: &ListNode{Val: 1, Next: &ListNode{Val: 3}},
			n:    2,
		},
		want: &ListNode{Val: 3},
	},
	{
		name: "case3",
		args: args{
			head: &ListNode{Val: 4},
			n:    1,
		},
		want: nil,
	},
	{
		name: "case4",
		args: args{
			head: nil,
			n:    0,
		},
		want: nil,
	},
}

func Test_removeNthFromEnd(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeNthFromEnd(tt.args.head, tt.args.n))
		})
	}
}

func Test_removeNthFromEndStack(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeNthFromEndStack(tt.args.head, tt.args.n))
		})
	}
}

func Test_removeNthFromEndDoublePtr(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeNthFromEndDoublePtr(tt.args.head, tt.args.n))
		})
	}
}

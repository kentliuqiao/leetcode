package removeelementsinlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	head *ListNode
	val  int
}

var tests = []struct {
	name string
	args args
	want *ListNode
}{
	{
		name: "case1",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}},
			val:  4,
		},
		want: &ListNode{Val: 2, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}},
	},
	{
		name: "case2",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}},
			val:  8,
		},
		want: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}},
	},
	{
		name: "case3",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}}},
			val:  2,
		},
		want: nil,
	},
	{
		name: "case4",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}},
			val:  1,
		},
		want: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}},
	},
	{
		name: "case5",
		args: args{
			head: nil,
			val:  1,
		},
		want: nil,
	},
}

func Test_removeElements(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeElements(tt.args.head, tt.args.val))
		})

	}
}

func Test_removeElementsRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeElementsRecursive(tt.args.head, tt.args.val))
		})

	}
}

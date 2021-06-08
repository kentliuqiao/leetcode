package detectcyclelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	head *ListNode
}

var (
	node1 = &ListNode{Val: 1}
	node2 = &ListNode{Val: 2}
	node3 = &ListNode{Val: 1}
	node4 = &ListNode{Val: 1}
	node5 = &ListNode{Val: 1}
)

var tests = []struct {
	name string
	args args
	want *ListNode
}{
	{
		name: "case1",
		args: args{
			head: node1,
		},
		want: node1,
	},
	{
		name: "case2",
		args: args{
			head: node3,
		},
		want: node4,
	},
	{
		name: "case3",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
		},
		want: nil,
	},
	{
		name: "case4",
		args: args{head: nil},
		want: nil,
	},
	{
		name: "case5",
		args: args{head: &ListNode{Val: 1}},
		want: nil,
	},
}

func Test_detectCycle(t *testing.T) {
	setUp()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, detectCycle(tt.args.head))
		})
	}
}

func setUp() {
	node1.Next = node2
	node2.Next = node1

	node3.Next = node4
	node4.Next = node5
	node5.Next = node4
}

func Test_detectCycleFastSlowPtr(t *testing.T) {
	setUp()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, detectCycleFastSlowPtr(tt.args.head))
		})
	}
}

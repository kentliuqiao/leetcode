package hascyclelinkedlist

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
	want bool
}{
	{
		name: "case1",
		args: args{
			head: node1,
		},
		want: true,
	},
	{
		name: "case4",
		args: args{
			head: node3,
		},
		want: true,
	},
	{
		name: "case2",
		args: args{
			head: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
		},
		want: false,
	},
	{
		name: "case3",
		args: args{head: nil},
		want: false,
	},
}

func Test_hasCycle(t *testing.T) {
	setUp()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hasCycle(tt.args.head))
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

func Test_hasCycleFastSlowPointer(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hasCycleFastSlowPointer(tt.args.head))
		})
	}
}

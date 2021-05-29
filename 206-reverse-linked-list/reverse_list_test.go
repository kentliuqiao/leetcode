package reverselinkedlist

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
		name: "case 1",
		args: args{
			head: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 6}}},
		},
		want: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4}}},
	},
	{
		name: "case 2",
		args: args{head: &ListNode{Val: 2}},
		want: &ListNode{Val: 2},
	},
	{
		name: "case 3",
		args: args{head: nil},
		want: nil,
	},
}

func Test_reverseList(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseList(tt.args.head))
		})
	}
}

func Test_reverseListRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseListRecursive(tt.args.head))
		})
	}
}

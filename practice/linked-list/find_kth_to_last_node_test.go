package linkedlist

import (
	"reflect"
	"testing"
)

func Test_findKthToLastNode(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{head: &ListNode{Val: 2, Next: &ListNode{4, &ListNode{9, &ListNode{11, nil}}}}, k: 2}, want: &ListNode{9, nil}},
		{name: "2", args: args{head: &ListNode{Val: 2}, k: 1}, want: &ListNode{2, nil}},
		{name: "3", args: args{head: nil, k: 0}, want: nil},
		{name: "4", args: args{head: &ListNode{Val: 2}, k: 2}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthToLastNode(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKthToLastNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

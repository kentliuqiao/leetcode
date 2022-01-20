package middleofthelinkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{head: &ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, &ListNode{10, nil}}}}}}, want: &ListNode{6, nil}},
		{name: "2", args: args{&ListNode{1, &ListNode{3, &ListNode{5, nil}}}}, want: &ListNode{3, nil}},
		{name: "3", args: args{&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, want: &ListNode{4, nil}},
		{name: "4", args: args{&ListNode{4, &ListNode{1, nil}}}, want: &ListNode{1, nil}},
		{name: "4", args: args{nil}, want: nil},
		{name: "5", args: args{&ListNode{1, nil}}, want: &ListNode{1, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package mergeksortedlists

import (
	"reflect"
	"testing"
)

var l1 []*ListNode = []*ListNode{
	{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: &ListNode{Val: 10}}}},
	{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8}}},
	{Val: 5, Next: &ListNode{Val: 11, Next: &ListNode{Val: 13}}},
}
var want1 *ListNode = &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, &ListNode{13, nil}}}}}}}}}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{lists: l1}, want: want1},
		{name: "2", args: args{lists: []*ListNode{{1, nil}, {9, nil}, {4, nil}, {7, nil}, {}}}, want: &ListNode{0, &ListNode{1, &ListNode{4, &ListNode{7, &ListNode{9, nil}}}}}},
		{name: "3", args: args{lists: []*ListNode{}}, want: nil},
		{name: "4", args: args{lists: []*ListNode{{1, nil}, {9, nil}, {1, nil}}}, want: &ListNode{1, &ListNode{1, &ListNode{9, nil}}}},
		{name: "5", args: args{lists: []*ListNode{{}}}, want: &ListNode{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

package mergeksortedlists

import (
	"reflect"
	"testing"
)

func Test_divideConquer(t *testing.T) {
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
			if got := divideConquer(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideConquer() = %v, want %v", got, tt.want)
			}
		})
	}
}

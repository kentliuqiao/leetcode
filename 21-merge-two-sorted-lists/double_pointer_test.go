package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func Test_doublePointer(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{
			list1: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}},
			list2: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doublePointer(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doublePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

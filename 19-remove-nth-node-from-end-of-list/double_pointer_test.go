package removenthnodefromendoflist

import (
	"reflect"
	"testing"
)

func Test_doublePointer(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{head: &ListNode{2, &ListNode{3, &ListNode{6, nil}}}, n: 2}, want: &ListNode{2, &ListNode{6, nil}}},
		{name: "2", args: args{head: &ListNode{3, &ListNode{6, nil}}, n: 2}, want: &ListNode{6, nil}},
		{name: "3", args: args{head: &ListNode{2, nil}, n: 1}, want: nil},
		{name: "4", args: args{head: nil, n: 0}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doublePointer(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doublePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

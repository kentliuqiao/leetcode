package findallanagramsinastring

import (
	"reflect"
	"testing"
)

func Test_slidingWindow(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "1", args: args{s: "cbaebabacd", p: "abc"}, want: []int{0, 6}},
		{name: "1", args: args{s: "abab", p: "ab"}, want: []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingWindow(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

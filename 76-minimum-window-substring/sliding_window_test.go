package minimumwindowsubstring

import "testing"

func Test_slidingWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1", args: args{s: "ADOBECODEBANC", t: "ABC"}, want: "BANC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

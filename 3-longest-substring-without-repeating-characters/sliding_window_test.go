package longestsubstringnorepeating

import "testing"

func Test_slidingWindow(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: "abcabcbb"}, want: 3},
		{name: "1", args: args{s: "bbbbb"}, want: 1},
		{name: "1", args: args{s: "pwwkew"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingWindow(tt.args.s); got != tt.want {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

package permutationinstring

import "testing"

func Test_slidingWindow(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s1: "ab", s2: "eidbaooo"}, want: true},
		{name: "1", args: args{s1: "ab", s2: "eidboaooo"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingWindow(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("slidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

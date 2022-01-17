package dppractice

import (
	"testing"
)

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	// TODO: Add test cases.
	{
		name: "1",
		args: args{1},
		want: 1,
	},
	{
		name: "2",
		args: args{2},
		want: 5,
	},
	{
		name: "3",
		args: args{4},
		want: 10,
	},
	{
		name: "4",
		args: args{10},
		want: 30,
	},
	{
		name: "5",
		args: args{9},
		want: 25,
	},
}

func Test_cutRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutRecursive(p, tt.args.n); got != tt.want {
				t.Errorf("cutRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cutTopDownMemo(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutTopDownMemo(p, tt.args.n); got != tt.want {
				t.Errorf("cutTopDownMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cutBottomUp(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cutBottomUp(p, tt.args.n); got != tt.want {
				t.Errorf("cutBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

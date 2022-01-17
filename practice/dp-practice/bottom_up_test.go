package dppractice

import (
	"testing"
)

func Test_bottomUpMemoFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
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
			want: 1,
		},
		{
			name: "3",
			args: args{3},
			want: 2,
		},
		{
			name: "4",
			args: args{9},
			want: 34,
		},
		{
			name: "5",
			args: args{4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bottomUpMemoFibonacci(tt.args.n); got != tt.want {
				t.Errorf("bottomUpMemoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bottomUpNoMemoFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
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
			want: 1,
		},
		{
			name: "3",
			args: args{3},
			want: 2,
		},
		{
			name: "4",
			args: args{9},
			want: 34,
		},
		{
			name: "5",
			args: args{4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bottomUpNoMemoFibonacci(tt.args.n); got != tt.want {
				t.Errorf("bottomUpNoMemoFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

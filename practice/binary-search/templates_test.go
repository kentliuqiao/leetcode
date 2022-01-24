package binarysearch

import (
	"testing"
)

func Test_leftBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{nums: []int{1, 3, 5, 7, 9}, target: 3}, want: 1},
		{name: "1", args: args{nums: []int{1, 3, 3, 7, 9}, target: 3}, want: 1},
		{name: "1", args: args{nums: []int{1, 3, 3, 3, 9}, target: 3}, want: 1},
		{name: "1", args: args{nums: []int{1, 1, 3, 3, 9}, target: 1}, want: 0},
		{name: "1", args: args{nums: []int{1, 3, 5, 7, 9}, target: 0}, want: -1},
		{name: "1", args: args{nums: []int{1, 3, 5, 7, 9}, target: 13}, want: -1},
		{name: "1", args: args{nums: []int{}, target: 1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("leftBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightBound(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{nums: []int{1, 3, 5, 7, 9}, target: 3}, want: 1},
		{name: "1", args: args{nums: []int{1, 3, 3, 7, 9}, target: 3}, want: 2},
		{name: "1", args: args{nums: []int{1, 3, 3, 3, 9}, target: 3}, want: 3},
		{name: "1", args: args{nums: []int{1, 1, 3, 3, 9}, target: 1}, want: 1},
		{name: "1", args: args{nums: []int{1, 3, 5, 7, 9}, target: 0}, want: -1},
		{name: "1", args: args{nums: []int{1, 3, 5, 7, 9}, target: 13}, want: -1},
		{name: "1", args: args{nums: []int{}, target: 1}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightBound(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("rightBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

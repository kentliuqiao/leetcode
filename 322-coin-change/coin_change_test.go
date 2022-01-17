package coinchange

import (
	"testing"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{coins: []int{1, 2, 5}, amount: 2}, want: 1},
		{name: "2", args: args{coins: []int{1, 2, 5}, amount: 3}, want: 2},
		{name: "3", args: args{coins: []int{1, 2, 5}, amount: 5}, want: 1},
		{name: "4", args: args{coins: []int{1, 2, 5}, amount: 7}, want: 2},
		{name: "5", args: args{coins: []int{1, 2, 5}, amount: 8}, want: 3},
		{name: "6", args: args{coins: []int{1, 2, 5}, amount: 0}, want: 0},
		{name: "7", args: args{coins: []int{1, 2, 5}, amount: -1}, want: -1},
		{name: "8", args: args{coins: []int{1, 2, 5}, amount: 10}, want: 2},
		{name: "9", args: args{coins: []int{2, 4}, amount: 3}, want: -1},
		{name: "9", args: args{coins: []int{}, amount: 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bottomUp(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{coins: []int{1, 2, 5}, amount: 2}, want: 1},
		{name: "2", args: args{coins: []int{1, 2, 5}, amount: 3}, want: 2},
		{name: "3", args: args{coins: []int{1, 2, 5}, amount: 5}, want: 1},
		{name: "4", args: args{coins: []int{1, 2, 5}, amount: 7}, want: 2},
		{name: "5", args: args{coins: []int{1, 2, 5}, amount: 8}, want: 3},
		{name: "6", args: args{coins: []int{1, 2, 5}, amount: 0}, want: 0},
		{name: "7", args: args{coins: []int{1, 2, 5}, amount: -1}, want: -1},
		{name: "8", args: args{coins: []int{1, 2, 5}, amount: 10}, want: 2},
		{name: "9", args: args{coins: []int{2, 4}, amount: 3}, want: -1},
		{name: "10", args: args{coins: []int{}, amount: 3}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bottomUp(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("bottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

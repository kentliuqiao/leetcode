package dppractice

import (
	"testing"
)

var coins []int = []int{1, 2, 5}

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
		{name: "1", args: args{coins: coins, amount: 2}, want: 1},
		{name: "2", args: args{coins: coins, amount: 3}, want: 2},
		{name: "3", args: args{coins: coins, amount: 5}, want: 1},
		{name: "4", args: args{coins: coins, amount: 7}, want: 2},
		{name: "5", args: args{coins: coins, amount: 8}, want: 3},
		{name: "6", args: args{coins: coins, amount: 0}, want: 0},
		{name: "7", args: args{coins: coins, amount: -1}, want: -1},
		{name: "8", args: args{coins: coins, amount: 10}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChangeWithMemo(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{coins: coins, amount: 2}, want: 1},
		{name: "2", args: args{coins: coins, amount: 3}, want: 2},
		{name: "3", args: args{coins: coins, amount: 5}, want: 1},
		{name: "4", args: args{coins: coins, amount: 7}, want: 2},
		{name: "5", args: args{coins: coins, amount: 8}, want: 3},
		{name: "6", args: args{coins: coins, amount: 0}, want: 0},
		{name: "7", args: args{coins: coins, amount: -1}, want: -1},
		{name: "8", args: args{coins: coins, amount: 10}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeWithMemo(coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangeWithMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChangeBottomUp(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{coins: coins, amount: 2}, want: 1},
		{name: "2", args: args{coins: coins, amount: 3}, want: 2},
		{name: "3", args: args{coins: coins, amount: 5}, want: 1},
		{name: "4", args: args{coins: coins, amount: 7}, want: 2},
		{name: "5", args: args{coins: coins, amount: 8}, want: 3},
		{name: "6", args: args{coins: coins, amount: 0}, want: 0},
		{name: "7", args: args{coins: coins, amount: -1}, want: -1},
		{name: "8", args: args{coins: coins, amount: 10}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeBottomUp(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChangeBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

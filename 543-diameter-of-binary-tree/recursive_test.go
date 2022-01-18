package diameterofbinarytree

import (
	"testing"
)

func Test_recursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{root: nil}, want: 0},
		{name: "2", args: args{root: &TreeNode{}}, want: 0},
		{name: "3", args: args{root: &TreeNode{Right: &TreeNode{}}}, want: 1},
		{name: "4", args: args{root: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{Left: &TreeNode{}}}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := recursive(tt.args.root); got != tt.want {
			// 	t.Errorf("recursive() = %v, want %v", got, tt.want)
			// }
			if got := divideConquer(tt.args.root); got != tt.want {
				t.Errorf("divideConquer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{root: nil}, want: 0},
		{name: "2", args: args{root: &TreeNode{}}, want: 1},
		{name: "3", args: args{root: &TreeNode{Right: &TreeNode{}}}, want: 2},
		{name: "4", args: args{root: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{Left: &TreeNode{}}}}, want: 3},
		{name: "4", args: args{root: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{Left: &TreeNode{Left: &TreeNode{}}}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

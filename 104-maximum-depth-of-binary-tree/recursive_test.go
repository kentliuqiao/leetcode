package maximumdepthofbinarytree

import "testing"

var tree *TreeNode = &TreeNode{
	Val:  3,
	Left: &TreeNode{Val: 9},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 9},
	},
}

func Test_recursive(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{root: tree}, want: 3},
		{name: "2", args: args{root: &TreeNode{}}, want: 1},
		{name: "3", args: args{root: &TreeNode{Right: &TreeNode{}}}, want: 2},
		{name: "4", args: args{root: nil}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recursive(tt.args.root); got != tt.want {
				t.Errorf("recursive() = %v, want %v", got, tt.want)
			}
			if got := traverse(tt.args.root); got != tt.want {
				t.Errorf("traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

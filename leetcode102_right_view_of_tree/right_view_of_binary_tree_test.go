package leetcode102_right_view_of_tree

import (
	"reflect"
	"testing"
)

func TestRightViewOfBinaryTree(t *testing.T) {
	tree := &TreeNode{
		val: 12,
		left: &TreeNode{
			val: 3,
			left: &TreeNode{
				val: 83,
				right: &TreeNode{
					val: 5,
				},
			},
		},
		right: &TreeNode{
			val: 4,
			right: &TreeNode{
				val: 0,
			},
		},
	}
	got := rightViewOfBinaryTree(tree)
	want := []int{12, 4, 0, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("rightViewOfBinaryTree want %v, got %v", want, got)
	}
}

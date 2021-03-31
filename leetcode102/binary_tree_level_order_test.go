package leetcode102

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
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
				left: &TreeNode{
					val: 9,
				},
			},
		},
	}
	want := [][]int{{12}, {3, 4}, {83, 0}, {5, 9}}
	got := levelOrder(tree)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("level order got %v, but want %v", got, want)
	}
}

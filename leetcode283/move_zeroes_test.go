package leetcode283

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
	}

	for _, c := range cases {
		MoveZeroes(c.nums)
		if !reflect.DeepEqual(c.nums, c.want) {
			t.Errorf("MoveZeroes got %v but want %v", c.nums, c.want)
		}
	}
}

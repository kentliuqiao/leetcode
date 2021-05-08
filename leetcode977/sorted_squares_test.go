package leetcode977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	nums []int
	want []int
}{
	{
		nums: []int{-4, -1, 0, 3, 10},
		want: []int{0, 1, 9, 16, 100},
	},
	{
		nums: []int{-7, -3, 2, 3, 11},
		want: []int{4, 9, 9, 49, 121},
	},
	{
		nums: []int{},
		want: []int{},
	},
	{
		nums: []int{-2},
		want: []int{4},
	},
}

func TestSortedSquares(t *testing.T) {
	for _, c := range cases {
		got := sortedSquares(c.nums)
		assert.Equal(t, c.want, got)
	}
}

func TestSortedSquaresOptimization_1(t *testing.T) {
	for _, c := range cases {
		got := sortedSquaresOptimization1(c.nums)
		assert.Equal(t, c.want, got)
	}
}

package leetcode209

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	target int
	nums   []int
	want   int
}{
	{
		target: 7,
		nums:   []int{2, 3, 1, 2, 4, 3},
		want:   2,
	},
	{
		target: 4,
		nums:   []int{1, 4, 4},
		want:   1,
	},
	{
		target: 11,
		nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		want:   0,
	},
	{
		target: 1,
		nums:   []int{},
		want:   0,
	},
}

func TestMinSubArrayLen(t *testing.T) {
	for _, c := range cases {
		got := minSubArrayLen(c.target, c.nums)
		assert.Equal(t, c.want, got)
	}
}

func TestMinSubArrayLenOptimization(t *testing.T) {
	for _, c := range cases {
		got := minSubArrayLenOptimization(c.target, c.nums)
		assert.Equal(t, c.want, got)
	}
}

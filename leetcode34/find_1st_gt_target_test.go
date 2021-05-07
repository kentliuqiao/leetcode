package leetcode34

import "testing"

var casesGT = []struct {
	target int
	nums   []int
	want   int
}{
	{
		target: 6,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   3,
	},
	{
		target: 12,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   -1,
	},
	{
		target: -4,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   0,
	},
	{
		target: -4,
		nums:   []int{},
		want:   -1,
	},
}

func TestFindFirstElementGreaterThanTarget(t *testing.T) {
	for _, c := range casesGT {
		got := FindFirstElementGreaterThanTarget(c.target, c.nums)
		if got != c.want {
			t.Errorf("FindFirstElementGreaterThanTarget want %d but got %d", c.want, got)
		}
	}
}

func TestFindFirstElementGTOrGTETarget_GT(t *testing.T) {
	for _, c := range casesGT {
		got := FindFirstElementGTOrGTETarget(c.target, c.nums, false)
		if got != c.want {
			t.Errorf("FindFirstElementGTOrGTETarget want %d but got %d", c.want, got)
		}
	}
}

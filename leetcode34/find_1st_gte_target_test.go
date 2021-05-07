package leetcode34

import "testing"

var casesGTE = []struct {
	target int
	nums   []int
	want   int
}{
	{
		target: 8,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   3,
	},
	{
		target: 4,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   3,
	},
	{
		target: -4,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   0,
	},
	{
		target: 12,
		nums:   []int{-1, 2, 3, 8, 10},
		want:   -1,
	},
	{
		target: -4,
		nums:   []int{},
		want:   -1,
	},
}

func TestFindFirstElementGreaterThanOrEqualTarget(t *testing.T) {

	for _, c := range casesGTE {
		got := FindFirstElementGreaterThanOrEqualTarget(c.target, c.nums)
		if got != c.want {
			t.Errorf("FindFirstElementGreaterThanTarget want %d but got %d", c.want, got)
		}
	}
}

func TestFindFirstElementGTOrGTETarget_GTE(t *testing.T) {
	for _, c := range casesGTE {
		got := FindFirstElementGTOrGTETarget(c.target, c.nums, true)
		if got != c.want {
			t.Errorf("FindFirstElementGTOrGTETarget want %d but got %d", c.want, got)
		}
	}
}

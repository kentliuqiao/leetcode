package leetcode34

import (
	"reflect"
	"testing"
)

var casesRange = []struct {
	target int
	nums   []int
	want   []int
}{
	{
		target: 8,
		nums:   []int{5, 7, 7, 8, 8, 10},
		want:   []int{3, 4},
	},
	{
		target: 6,
		nums:   []int{5, 7, 7, 8, 8, 10},
		want:   []int{-1, -1},
	},
	{
		target: 8,
		nums:   []int{},
		want:   []int{-1, -1},
	},
	{
		target: 5,
		nums:   []int{5, 7, 7, 8, 8, 10},
		want:   []int{0, 0},
	},
}

func TestSearcgRange(t *testing.T) {

	for _, c := range casesRange {
		got := SearchRange(c.target, c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SearchRange want %v but got %v", c.want, got)
		}
	}
}

func TestSearcgRangeOptimization(t *testing.T) {
	for _, c := range casesRange {
		got := SearchRangeOptmization(c.target, c.nums)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("SearchRangeOptmization want %d but got %d", c.want, got)
		}
	}
}

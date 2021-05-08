package leetcode844

import "testing"

var cases = []struct {
	s, t string
	want bool
}{
	{
		s:    "ab#c",
		t:    "ad#c",
		want: true,
	},
	{
		s:    "ab##",
		t:    "c#d#",
		want: true,
	},
	{
		s:    "a##c",
		t:    "#a#c",
		want: true,
	},
	{
		s:    "a#c",
		t:    "b",
		want: false,
	},
}

func TestBackSpaceCompare(t *testing.T) {

	for _, c := range cases {
		got := backSpaceCompare(c.s, c.t)
		if got != c.want {
			t.Errorf("backSpaceCompare want %t but got %t, s: %s, t: %s", got, c.want, c.s, c.t)
		}
	}
}

func TestBackSpaceCompareOptimization(t *testing.T) {
	for _, c := range cases {
		got := backSpaceCompareOptimization(c.s, c.t)
		if got != c.want {
			t.Errorf("backSpaceCompare want %t but got %t, s: %s, t: %s", got, c.want, c.s, c.t)
		}
	}
}

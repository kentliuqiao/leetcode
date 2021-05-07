package leetcode35

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		target int
		nums   []int
		want   int
	}{
		{
			target: 5,
			nums:   []int{1, 3, 5, 6},
			want:   2,
		},
		{
			target: 2,
			nums:   []int{1, 3, 5, 6},
			want:   1,
		},
		{
			target: 0,
			nums:   []int{1, 3, 5, 6},
			want:   0,
		},
		{
			target: 7,
			nums:   []int{1, 3, 5, 6},
			want:   4,
		},
	}

	for _, c := range cases {
		got := SearchInsert(c.target, c.nums)
		if got != c.want {
			t.Errorf("SearchInsert want %d, but got %d", c.want, got)
		}
	}
}

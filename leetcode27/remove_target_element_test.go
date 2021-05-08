package leetcode27

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		target int
		nums   []int
		want   struct {
			cnt  int
			nums []int
		}
	}{
		{
			target: 3,
			nums:   []int{3, 2, 2, 3},
			want: struct {
				cnt  int
				nums []int
			}{
				cnt:  2,
				nums: []int{2, 2, 3, 0},
			},
		},
		{
			target: 4,
			nums:   []int{3, 2, 2, 3},
			want: struct {
				cnt  int
				nums []int
			}{
				cnt:  0,
				nums: []int{},
			},
		},
	}

	for _, c := range cases {
		cnt := RemoveElement(c.target, c.nums)
		if cnt != c.want.cnt || !_assertArrayEq(cnt, c.nums, c.want.nums) {
			t.Errorf("RemoveElement got cnt %d and nums %v but want %d and nums %v", cnt, c.nums, c.want.cnt, c.want.nums)
		}
	}
}

func _assertArrayEq(cnt int, x1, x2 []int) bool {
	for i := 0; i < cnt; i++ {
		if x1[i] != x2[i] {
			return false
		}
	}

	return true
}

package leetcode26

import "testing"

func TestRemoveRepeatedElement(t *testing.T) {
	cases := []struct {
		nums []int
		want struct {
			nums []int
			cnt  int
		}
	}{
		{
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: struct {
				nums []int
				cnt  int
			}{
				nums: []int{0, 1, 2, 3, 4, 0, 0},
				cnt:  5,
			},
		},
		{
			nums: []int{1, 1, 2},
			want: struct {
				nums []int
				cnt  int
			}{
				nums: []int{1, 2},
				cnt:  2,
			},
		},
	}

	for _, c := range cases {
		got := RemoveRepeatedElement(c.nums)
		if got != c.want.cnt || !_assertSliceEqual(got, c.nums, c.want.nums) {
			t.Errorf("RemoveRepeatedElement got cnt %d nums %v but want cnt %d nums %v", got, c.nums, c.want.cnt, c.want.nums)
		}
	}
}

func _assertSliceEqual(cnt int, x1, x2 []int) bool {
	for i := 0; i < cnt; i++ {
		if x1[i] != x2[i] {
			return false
		}
	}

	return true
}

package leetcode34

func SearchRange(target int, nums []int) []int {
	pivot, left, right := 0, 0, len(nums)-1
	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] == target {
			i := pivot - 1
			for ; i >= 0 && nums[i] == target; i-- {
			}
			j := pivot + 1
			for ; j < len(nums) && nums[j] == target; j++ {
			}
			return []int{i + 1, j - 1}
		}
		if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return []int{-1, -1}
}

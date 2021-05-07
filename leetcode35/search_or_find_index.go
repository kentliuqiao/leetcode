package leetcode35

func SearchInsert(target int, nums []int) int {
	pivot, left, right := 0, 0, len(nums)-1
	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		}
		if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return right + 1
}

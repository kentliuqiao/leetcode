package searchmininrevertedarray

import "math"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt64
	}
	min := nums[0]
	left, right := 1, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] > min {
			left = mid + 1
		} else {
			min = nums[mid]
			right = mid - 1
		}
	}
	return min
}

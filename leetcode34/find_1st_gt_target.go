package leetcode34

// FindFirstElementGreaterThanTarget find the first element/index in a sorted array
// that is greater than target using binary search
func FindFirstElementGreaterThanTarget(target int, nums []int) int {
	pivot, left, right, ans := 0, 0, len(nums)-1, len(nums)

	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] > target {
			ans = pivot
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	if ans == len(nums) {
		ans = -1
	}

	return ans
}

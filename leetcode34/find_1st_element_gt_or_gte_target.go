package leetcode34

// FindFirstElementGTOrGTETarget behaves like FindFirstElementGreaterThanTarget when lower is false
// and behaves like FindFirstElementGreaterThanOrEqualTarget when lower is true
func FindFirstElementGTOrGTETarget(target int, nums []int, lower bool) int {
	pivot, left, right, ans := 0, 0, len(nums)-1, len(nums)
	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] > target || (lower && nums[pivot] >= target) {
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

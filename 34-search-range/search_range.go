package searchrange

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}
	leftIdx := binarySearch(nums, target, true)
	rightIdx := binarySearch(nums, target, false) - 1
	if leftIdx <= rightIdx && rightIdx < len(nums) && target == nums[leftIdx] && target == nums[rightIdx] {
		return []int{leftIdx, rightIdx}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, target int, lower bool) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (lower && nums[mid] == target) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

package leetcode34

func SearchRangeOptmization(target int, nums []int) []int {
	left := FindFirstElementGTOrGTETarget(target, nums, true)
	right := FindFirstElementGTOrGTETarget(target, nums, false) - 1
	if left <= right && left >= 0 && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}

	return []int{-1, -1}
}

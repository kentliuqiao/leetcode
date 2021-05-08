package leetcode27

func RemoveElement(target int, nums []int) int {
	left, right, cnt := 0, 0, 0

	for right < len(nums) {
		if nums[right] != target {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		} else {
			cnt++
		}
		right++
	}

	return cnt
}

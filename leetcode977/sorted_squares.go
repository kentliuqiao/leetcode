package leetcode977

import "sort"

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	// time: O(nlogn)
	// space: O(n)
	sort.Ints(nums)

	return nums
}

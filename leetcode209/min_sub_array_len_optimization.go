package leetcode209

import "math"

// 滑动窗口
func minSubArrayLenOptimization(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sum := 0
	ans := math.MaxInt32
	start, end := 0, 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

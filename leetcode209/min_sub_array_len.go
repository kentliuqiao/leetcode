package leetcode209

func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sumSubArray(nums[i:j]) >= target {
				if j-i < min {
					min = j - i
				}
				break
			}
		}
	}
	if min > len(nums) {
		min = 0
	}

	return min
}

func sumSubArray(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

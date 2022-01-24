package binarysearch

/**
===================================================================
                        二分查找基本模版
===================================================================

技巧：
1. 不要出现 else，而是把所有情况用 else if 写清楚，这样可以清楚地展现所有细节。

func binarySearch(nums []int, target int) int {
	left, right := 0, ...
	for ... {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			...
		} else if nums[mid] > target {
			right = ...
		} else if nums[mid] < target {
			left = ...
		}
	}
	return ...
}

// 寻找一个数
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// 寻找左侧边界
// 注意：当target比所有元素都小时，需要返回0；当target比所有元素都大时，需要返回len(nums)。
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums) // 因为有可能取到len(nums)，所以right要初始化为len(nums)
	for left < right { // 无论target是否存在于nums中，最后都要返回一个index，这个index一定是最后left和right相遇的位置，因此使用 < 。
		mid := left + (right-left)/2
		if nums[mid] == target { // 因为需要找到左边界，所以在找到匹配元素时，还得不断向左探索
			right = mid
		} else if nums[mid] < target { // 当target比mid处元素大时，最终答案必不可能取到此位置，因此left=mid+1
			left = mid + 1
		} else if nums[mid] > target { // 当target比mid处元素小时，最终答案是有可能取到mid的，因此right=mid
			right = mid
		}
	}
	return left
}

// 寻找左侧边界
func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 寻找右边界
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left == len(nums) {
		return -1
	}
	return left - 1
}

// 寻找右边界
func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target { // 检查右边界情况
		return -1
	}
	return right
}


*/

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target { // 检查右边界情况
		return -1
	}
	return right
}

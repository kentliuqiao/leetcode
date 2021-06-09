package searchinvertedarray

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		// 以mid为分隔，将数组分为左右两个部分，则其中一个部分一定有序，另外一个部分有可能有序！
		if nums[left] <= nums[mid] {
			// 数组左半部分有序
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 数组右半部分有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

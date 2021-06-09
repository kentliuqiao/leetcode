package searchbidimensionalarray

func searchMatrix(nums [][]int, target int) bool {
	if len(nums) == 0 || len(nums[0]) == 0 {
		return false
	}
	rowIdx := searchFirstColumn(nums, target)
	if rowIdx < 0 {
		return false
	}
	return searchRow(nums[rowIdx], target)
}

func searchFirstColumn(nums [][]int, target int) int {
	low, high := -1, len(nums)-1
	for low < high {
		mid := low + (high-low+1)/2
		if nums[mid][0] <= target {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func searchRow(row []int, target int) bool {
	left, right := 0, len(row)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target == row[mid] {
			return true
		}
		if target < row[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

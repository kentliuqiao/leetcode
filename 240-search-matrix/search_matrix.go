package searchmatrixii

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	colIdx := searchRow(matrix, target, 0)
	if colIdx < 0 {
		return false
	}
	rowIdx := searchColumn(matrix, target, colIdx)
	if rowIdx < 0 {
		return false
	}
	if target != matrix[rowIdx][colIdx] {
		colIdx = searchRow(matrix, target, rowIdx)
	}
	if colIdx < 0 {
		return false
	}
	return matrix[rowIdx][colIdx] == target
}

func searchRow(matrix [][]int, target, row int) int {
	left, right := -1, len(matrix[row])-1
	for left < right {
		mid := left + (right-left+1)>>1
		if matrix[row][mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchColumn(matrix [][]int, target, column int) int {
	low, high := -1, len(matrix)-1
	for low < high {
		mid := low + (high-low+1)>>1
		if matrix[mid][column] <= target {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func searchMatrixV2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, col := len(matrix)-1, 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}

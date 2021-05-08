package leetcode977

func sortedSquaresOptimization1(nums []int) []int {
	n := len(nums)
	m := n - 1
	res := make([]int, n)
	i, j := 0, n-1
	for i <= j {
		head := square(nums[i])
		tail := square(nums[j])
		if head < tail {
			res[m] = tail
			m--
			j--
		} else {
			res[m] = head
			m--
			i++
		}
	}

	return res
}

func square(num int) int {
	return num * num
}

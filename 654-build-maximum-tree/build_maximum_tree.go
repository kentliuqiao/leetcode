package buildmaximumtree

type TreeNode struct {
	val         int
	left, right *TreeNode
}

// Note: all number in nums are between 0 and 1000

func builMaximumTree(nums []int) *TreeNode {
	idx, max := findMax(nums)
	if idx == -1 || max == -1 {
		return nil
	}
	root := &TreeNode{val: max}
	root.left = builMaximumTree(nums[:idx])
	root.right = builMaximumTree(nums[idx+1:])
	return root
}

func findMax(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}
	i, max := 0, nums[0]
	for k, n := range nums {
		if n > max {
			max = n
			i = k
		}
	}
	return i, max
}

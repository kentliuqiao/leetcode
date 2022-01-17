package permutations

var res = [][]int{}

// 回溯算法计算全排列
func permute(nums []int) [][]int {
	defer func() {
		res = [][]int{}
	}()
	track := []int{}
	backTrack(nums, track)
	return res
}

func backTrack(nums, track []int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := range nums {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		backTrack(nums, track)
		track = track[:len(track)-1]
	}
}

func contains(track []int, n int) bool {
	for _, v := range track {
		if v == n {
			return true
		}
	}
	return false
}

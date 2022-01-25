package longestsubstringnorepeating

func slidingWindow(s string) int {
	seen := map[byte]int{}

	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++ // 不断扩大窗口
		seen[c]++
		for seen[c] > 1 { // c 字符有重复，需要减小窗口
			d := s[left]
			left++
			seen[d]--
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

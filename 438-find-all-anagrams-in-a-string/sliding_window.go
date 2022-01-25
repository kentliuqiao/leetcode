package findallanagramsinastring

func slidingWindow(s, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range []byte(p) {
		need[c]++
	}

	valid := 0
	left, right := 0, 0
	res := []int{}

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left == len(p) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

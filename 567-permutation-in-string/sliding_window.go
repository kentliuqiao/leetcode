package permutationinstring

func slidingWindow(s1, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range []byte(s1) {
		need[c]++
	}

	valid := 0
	left, right := 0, 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

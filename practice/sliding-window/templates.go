package slidingwindow

/**
========================================================================
						滑动窗口模版
========================================================================

func slidingWindow(s, t string) {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s) {
		// c 是即将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内的一系列更新
		// ...

		// debug 输出的位置
		fmt.Printf("window: [%d %d)\n", left, right)

		for (window needs shrink) {
			// d 是即将移出窗口的字符
			d := s[left]
			// 坐移窗口
			left++
			// 进行窗口内的一系列更新
			// ...
		}
	}
}

*/

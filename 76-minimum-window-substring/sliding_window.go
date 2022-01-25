package minimumwindowsubstring

import (
	"math"
)

func slidingWindow(s, t string) string {
	// need 记录需要匹配的子串字符
	// window 记录当前窗口内的字符
	need, window := make(map[byte]int), make(map[byte]int)
	for _, c := range []byte(t) {
		need[c]++
	}

	// 记录满足要求的子串起始位置以及子串长度
	start, length := 0, math.MaxInt64
	// 表示window窗口内满足要求的字符个数，当 valid == len(need) 时，说明已经满足了要求
	valid := 0

	// 表示窗口左右边界，左闭右开 [left, right)
	left, right := 0, 0
	for right < len(s) {
		// 即将移入窗口的字符
		c := s[right]
		// 窗口右侧右移
		right++
		// 进行窗口内一系列数据的更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] { // 这个字符已经满足要求
				valid++
			}
		}

		// 判断窗口左侧是否需要收缩
		for valid == len(need) {
			// 更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// 即将移出窗口的字符
			d := s[left]
			// 窗口左侧右移
			left++
			// 进行窗口内一系列数据的更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] { // d 移出后，窗口内字符 d 不再满足要求，因此valid需要减1
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt64 {
		return ""
	}
	return s[start : start+length]
}

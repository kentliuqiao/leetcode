package leetcode844

func backSpaceCompare(s, t string) bool {
	return buildString(s) == buildString(t)
}

func buildString(str string) string {
	res := []byte{}
	for i := range str {
		if str[i] != '#' {
			res = append(res, str[i])
		} else if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}

	return string(res)
}

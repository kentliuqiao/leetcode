package openthelock

// 24 ms	6.7 MB
func biDirectionalBFS(deadends []string, target string) int {
	visited := map[string]struct{}{}
	for _, v := range deadends {
		visited[v] = struct{}{}
	}
	q1, q2 := map[string]struct{}{"0000": {}}, map[string]struct{}{target: {}}
	step := 0

	for len(q1) > 0 && len(q2) > 0 {
		temp := map[string]struct{}{}
		for curr := range q1 {
			if _, ok := visited[curr]; ok {
				continue
			}
			if _, ok := q2[curr]; ok {
				return step
			}
			visited[curr] = struct{}{}
			for i := 0; i < 4; i++ {
				plus := plusOne(curr, i)
				if _, ok := visited[plus]; !ok {
					temp[plus] = struct{}{}
				}
				minus := minusOne(curr, i)
				if _, ok := visited[minus]; !ok {
					temp[minus] = struct{}{}
				}
			}
		}
		step++
		q1 = q2
		q2 = temp
	}
	return -1
}

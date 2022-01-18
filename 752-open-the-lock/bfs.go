package openthelock

// 	92 ms	7.4 MB
func bfs(deadends []string, target string) int {
	if target == "" {
		return -1
	}
	deadendsMap := make(map[string]struct{})
	for _, v := range deadends {
		deadendsMap[v] = struct{}{}
	}
	visited := make(map[string]struct{})
	q := []string{"0000"}
	visited["0000"] = struct{}{}
	steps := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[i]
			if _, ok := deadendsMap[curr]; ok {
				continue
			}
			if curr == target {
				return steps
			}
			for j := 0; j < 4; j++ {
				plus := plusOne(curr, j)
				if _, ok := visited[plus]; !ok {
					q = append(q, plus)
					visited[plus] = struct{}{}

				}
				minus := minusOne(curr, j)
				if _, ok := visited[minus]; !ok {
					q = append(q, minus)
					visited[minus] = struct{}{}
				}
			}
		}
		steps++
		q = q[size:]
	}
	return -1
}

func plusOne(seq string, index int) string {
	bytes := []byte(seq)
	if bytes[index] == '9' {
		bytes[index] = '0'
	} else {
		bytes[index] += 1
	}
	return string(bytes)
}

func minusOne(seq string, index int) string {
	bytes := []byte(seq)
	if bytes[index] == '0' {
		bytes[index] = '9'
	} else {
		bytes[index] -= 1
	}
	return string(bytes)
}

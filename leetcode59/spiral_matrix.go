package leetcode59

type pair struct{ x, y int }

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateSpiralMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	if n == 0 {
		return [][]int{}
	}

	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx]
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] != 0 {
			// change direction
			dirIdx = (dirIdx + 1) % 4
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}

	return matrix
}

package countsubislands

func countSubIslands(grid1, grid2 [][]int) int {
	res := 0
	m, n := len(grid2), len(grid2[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && floodFill(grid2, grid1, i, j) {
				res++
			}
		}
	}
	return res
}

func floodFill(grid, grid1 [][]int, i, j int) bool {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return true
	}
	if grid[i][j] == 0 {
		return true
	}
	if grid1[i][j] != grid[i][j] {
		return false
	}
	grid[i][j] = 0
	up := floodFill(grid, grid1, i-1, j)
	down := floodFill(grid, grid1, i+1, j)
	left := floodFill(grid, grid1, i, j-1)
	right := floodFill(grid, grid1, i, j+1)
	return up && down && left && right
}

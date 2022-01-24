package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res = max(res, floodFill(grid, i, j))
			}
		}
	}
	return res
}

func floodFill(grid [][]int, i, j int) int {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}
	if grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	left := floodFill(grid, i, j-1)
	right := floodFill(grid, i, j+1)
	up := floodFill(grid, i-1, j)
	down := floodFill(grid, i+1, j)
	return 1 + left + right + up + down
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

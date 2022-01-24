package numberofenclaves

func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for j := 0; j < n; j++ {
		floodFill(grid, 0, j)
		floodFill(grid, m-1, j)
	}
	for i := 0; i < m; i++ {
		floodFill(grid, i, 0)
		floodFill(grid, i, n-1)
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}

func floodFill(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	floodFill(grid, i-1, j)
	floodFill(grid, i+1, j)
	floodFill(grid, i, j-1)
	floodFill(grid, i, j+1)
}

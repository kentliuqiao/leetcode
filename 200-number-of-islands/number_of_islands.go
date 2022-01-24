package numberofislands

func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				floodFill(grid, i, j)
			}
		}
	}
	return res
}

// floodFill 将一个岛屿节点（1）周围（上下左右）的所有相连的岛屿淹没（0）
// 以让相连的所有的岛屿只会被计数一次
func floodFill(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	floodFill(grid, i-1, j)
	floodFill(grid, i+1, j)
	floodFill(grid, i, j-1)
	floodFill(grid, i, j+1)
}

package numberofclosedislands

// closedIsland
// 思路：首先淹没上下左右与陆地接壤的岛屿
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for j := 0; j < n; j++ {
		dfs(grid, 0, j)   // 淹没最上面接壤陆地的岛屿
		dfs(grid, m-1, j) // 淹没最下面接壤陆地的岛屿
	}
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)   // 淹没最左面接壤陆地的岛屿
		dfs(grid, i, n-1) // 淹没最右面接壤陆地的岛屿
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

package twodimensionarray

// dfs 遍历二维数组模版
func dfs(grid [][]int, i, j int, visited [][]bool) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if visited[i][j] {
		return
	}
	// 进入节点（i, j）
	visited[i][j] = true

	dfs(grid, i-1, j, visited) // 上
	dfs(grid, i+1, j, visited) // 下
	dfs(grid, i, j-1, visited) // 左
	dfs(grid, i, j+1, visited) // 右

	// 离开节点
}

// 方向数组，代表需要遍历的所有方向
// 在这里，分别代表上下左右四个方向
var dirs [][]int = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 使用方向数组dfs遍历二维数组
// 本质上和上面一样，只是使用循环来处理上下左右四个方向的遍历
func dfsWithDirArray(grid [][]int, i, j int, visited [][]bool) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n {
		return
	}
	if visited[i][j] {
		return
	}
	// 进入节点（i, j）
	visited[i][j] = true

	for _, dir := range dirs {
		nextI := i + dir[0]
		nextJ := i + dir[1]
		dfs(grid, nextI, nextJ, visited)
	}

	// 离开节点
}

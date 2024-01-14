package dfs

func closedIsland(grid [][]int) int {
	m := len(grid) - 1
	n := len(grid[0]) - 1
	for i := 0; i <= m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n)
	}

	for j := 0; j <= n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m, j)
	}

	result := 0
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] == 0 {
				result++
				dfs(grid, i, j)
			}
		}
	}
	return result
}

func dfs(grid [][]int, i, j int) {
	m := len(grid) - 1
	n := len(grid[0]) - 1
	if i < 0 || j < 0 || i > m || j > n {
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

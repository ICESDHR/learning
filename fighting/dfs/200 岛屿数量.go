package dfs

func numIslands(grid [][]byte) int {
	m := len(grid) - 1
	n := len(grid[0]) - 1
	result := 0
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if grid[i][j] == '1' {
				result++
				dfs(grid, i, j)
			}
		}
	}
	return result
}

func dfs(grid [][]byte, i, j int) {
	m := len(grid) - 1
	n := len(grid[0]) - 1
	if i < 0 || j < 0 || i > m || j > n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

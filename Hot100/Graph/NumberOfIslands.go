package Graph

var (
	visited   [][]bool
	direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

func numIslands(grid [][]byte) int {
	visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				ans++
				dfs(i, j, &grid)
			}
		}
	}
	return ans
}

func dfs(i, j int, grid *[][]byte) {
	visited[i][j] = true
	row := len(*grid)
	col := len((*grid)[0])
	for _, d := range direction {
		x := i + d[0]
		y := j + d[1]
		if x < 0 || x >= row || y < 0 || y >= col || visited[x][y] || (*grid)[x][y] == '0' {
			continue
		}
		dfs(x, y, grid)
	}
	return
}

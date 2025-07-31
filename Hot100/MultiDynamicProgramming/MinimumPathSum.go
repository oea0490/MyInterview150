package MultiDynamicProgramming

func minPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	// 定义dp[i][j]为到达第i行第j列的最小路径和
	// 状态转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// 初始状态：dp[0][0] = grid[0][0], dp[0][j] = dp[0][j-1] + grid[0][j], dp[i][0] = dp[i-1][0] + grid[i][0]
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	// 初始化
	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 状态转移
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[row-1][col-1]
}

package MultiDynamicProgramming

func uniquePaths(m int, n int) int {
	// 定义dp[i][j]为到达第i行第j列的路径数
	// 状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 初始状态：dp[0][j] = 1, dp[i][0] = 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 状态转移
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

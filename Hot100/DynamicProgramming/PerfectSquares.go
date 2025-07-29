package DynamicProgramming

func numSquares(n int) int {
	// 定义dp[i]为和为i的完全平方数的最少数量
	// 状态转移方程：dp[i] = min(dp[i], dp[i-j*j]+1)
	// 初始化：dp[0] = 0, dp[i*i] = i
	dp := make([]int, n+1)
	// 初始化
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for i := 1; i*i <= n; i++ {
		dp[i*i] = 1
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

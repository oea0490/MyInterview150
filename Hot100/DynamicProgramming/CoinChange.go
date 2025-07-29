package DynamicProgramming

func coinChange(coins []int, amount int) int {
	// 定义dp[i][j]为使用前i种面额的硬币凑成j金额的最少硬币数
	// 状态转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i]]+1)
	// 初始化：dp[i][0] = 0, dp[0][j] = 0x3f3f3f3f
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	// 初始化
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}
	for j := 1; j <= amount; j++ {
		dp[0][j] = 0x3f3f3f3f
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[n][amount] == 0x3f3f3f3f {
		return -1
	}
	return dp[n][amount]
}

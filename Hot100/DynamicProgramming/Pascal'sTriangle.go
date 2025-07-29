package DynamicProgramming

func generate(numRows int) [][]int {
	// 定义dp[i][j]为第i行第j列的杨辉三角数字
	// 状态转移方程：dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
	// 初始化：dp[0][0] = 1, dp[i][0] = 1, dp[i][i] = 1
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, numRows)
	}
	// 初始化
	dp[0][0] = 1
	for i := 1; i < numRows; i++ {
		dp[i][0] = 1
		dp[i][i] = 1
	}
	// 状态转移
	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	// 输出
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			ans[i][j] = dp[i][j]
		}
	}
	return ans
}

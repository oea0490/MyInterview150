package MultiDynamicProgramming

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	// 定义dp[i][j]为word1[:i]和word2[:j]的最小编辑距离
	// 状态转移方程：if word1[i] == word2[j] { dp[i][j] = dp[i-1][j-1] }
	//			   else { dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1 }
	// 初始状态：dp[0][j] = j, dp[i][0] = i
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	// 初始化
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}
	// 状态转移
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len1][len2]
}

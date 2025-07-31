package MultiDynamicProgramming

import "slices"

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	// 定义dp[i][j]为text1[:i]和text2[:j]的最长公共子序列长度
	// 状态转移方程：if text1[i] == text2[j] { dp[i][j] = dp[i-1][j-1] + 1 }
	//			   else { dp[i][j] = max(dp[i-1][j], dp[i][j-1]) }
	// 初始状态：dp[0][j] = 0, dp[i][0] = 0
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	// 初始化
	// 状态转移
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = slices.Max([]int{dp[i-1][j], dp[i][j-1]})
			}
		}
	}
	return dp[len1][len2]
}

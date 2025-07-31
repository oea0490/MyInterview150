package MultiDynamicProgramming

func longestPalindrome(s string) string {
	n := len(s)
	maxLen := 1
	start := 0
	// 定义dp[i][j]为s[i:j+1]是否为回文串
	// 状态转移方程：dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	// 初始状态：dp[i][i] = true, dp[i][i+1] = s[i] == s[i+1]
	dp := make([][]bool, n)
	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	// 状态转移
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if i+1 <= j-1 {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j]
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

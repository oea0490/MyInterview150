package DynamicProgramming

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	maxLen := 0
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
		maxLen = max(maxLen, len(word))
	}
	// 定义dp[i]为s[:i]是否可以被拆分成wordDict中的单词
	// 状态转移方程：dp[i] = dp[j] && s[j:i] in wordDict
	// 初始化：dp[0] = true
	dp := make([]bool, n+1)
	// 初始化
	dp[0] = true
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] = dp[j] && dict[s[j:i]]
			if dp[i] {
				break
			}
		}
	}
	return dp[n]
}

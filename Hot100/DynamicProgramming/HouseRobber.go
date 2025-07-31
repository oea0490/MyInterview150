package DynamicProgramming

func rob(nums []int) int {
	n := len(nums)
	if n < 2 {
		return nums[0]
	}
	// 定义dp[i]为偷到第i家时的最大金额
	// 状态转移方程：dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	// 初始化：dp[0] = nums[0], dp[1] = max(nums[0], nums[1])
	dp := make([]int, n)
	// 初始化
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	// 状态转移
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

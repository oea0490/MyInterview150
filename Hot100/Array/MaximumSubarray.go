package Array

func MaxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

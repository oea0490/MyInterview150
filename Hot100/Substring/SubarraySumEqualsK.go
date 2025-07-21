package Substring

func SubarraySum(nums []int, k int) int {
	preSumToCount := make(map[int]int)
	preSumToCount[0] = 1
	preSum := 0
	count := 0
	for _, num := range nums {
		preSum += num
		if _, ok := preSumToCount[preSum-k]; ok {
			count += preSumToCount[preSum-k]
		}
		preSumToCount[preSum]++
	}
	return count
}

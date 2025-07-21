package Array

func productExceptSelf(nums []int) []int {
	preProduct := make([]int, len(nums))
	postProduct := make([]int, len(nums))
	preProduct[0] = 1
	postProduct[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		preProduct[i] = preProduct[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		postProduct[i] = postProduct[i+1] * nums[i+1]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = preProduct[i] * postProduct[i]
	}
	return ans
}

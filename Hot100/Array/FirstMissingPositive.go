package Array

func FirstMissingPositive(nums []int) int {
	len := len(nums)
	for i := 0; i < len; i++ {
		for nums[i] > 0 && nums[i] <= len && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < len; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len + 1
}

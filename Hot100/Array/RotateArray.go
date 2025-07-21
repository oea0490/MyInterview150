package Array

func Rotate(nums []int, k int) {
	k %= len(nums)
	pre := nums[:len(nums)-k]
	last := nums[len(nums)-k:]
	copy(nums, append(last, pre...))
}

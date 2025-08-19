package BinarySearch

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}

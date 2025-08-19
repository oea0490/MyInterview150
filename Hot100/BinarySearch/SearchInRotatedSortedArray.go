package BinarySearch

func search(nums []int, target int) int {
	n := len(nums)
	// 第一次二分查找,先找到旋转点
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	rotate := left
	// 第二次二分查找,找到目标值
	left, right = rotate, rotate+n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid%n] > target {
			right = mid - 1
		} else if nums[mid%n] < target {
			left = mid + 1
		} else {
			return mid % n
		}
	}
	return -1
}

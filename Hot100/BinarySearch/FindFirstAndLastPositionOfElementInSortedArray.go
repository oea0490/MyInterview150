package BinarySearch

func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	ans[0] = findFirst(nums, target)
	ans[1] = findFirst(nums, target+1) - 1
	if ans[0] <= ans[1] {
		return ans
	}
	return []int{-1, -1}
}

func findFirst(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

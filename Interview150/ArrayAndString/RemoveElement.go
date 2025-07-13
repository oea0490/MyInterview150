package ArrayAndString

func RemoveElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

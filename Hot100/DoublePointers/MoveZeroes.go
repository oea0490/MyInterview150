package DoublePointers

func moveZeroes(nums []int) {
	point0 := 0
	point1 := 0
	for point1 < len(nums) {
		if nums[point1] != 0 {
			nums[point0], nums[point1] = nums[point1], nums[point0]
			point0++
		}
		point1++
	}
}

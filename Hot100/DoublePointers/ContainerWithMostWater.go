package DoublePointers

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		curArea := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, curArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

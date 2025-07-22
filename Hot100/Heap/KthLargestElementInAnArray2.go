package Heap

import "math/rand"

func FindKthLargestByQuickSelect(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	pivotIndex := rand.Intn(right-left+1) + left
	pivot := nums[pivotIndex]
	i, j := left-1, right+1
	for i < j {
		for {
			i++
			if nums[i] <= pivot {
				break
			}
		}
		for {
			j--
			if nums[j] >= pivot {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k > j {
		return quickSelect(nums, j+1, right, k)
	} else {
		return quickSelect(nums, left, j, k)
	}
}

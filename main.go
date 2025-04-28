package main

import (
	"Leetcode/Interview150/ArrayAndString"
	"fmt"
)

func main() {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	nums2 := []int{2, 4, 6}
	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)
	ArrayAndString.Merge(nums1, len(nums1)-len(nums2), nums2, len(nums2))
	fmt.Println("merged:", nums1)
}

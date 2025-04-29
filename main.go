package main

import (
	"Leetcode/Interview150/ArrayAndString"
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("before", nums)
	ArrayAndString.RemoveElement(nums, val)
	fmt.Println("after", nums)
}

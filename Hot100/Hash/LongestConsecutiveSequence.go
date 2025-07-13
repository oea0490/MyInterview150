package Hash

import "sort"

func LongestConsecutive(nums []int) int {
	sort.Ints(nums)
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	maxLen := 0
	for num, _ := range numSet {
		if !numSet[num-1] {
			curNum := num
			curLen := 1
			for numSet[curNum+1] {
				curNum += 1
				curLen += 1
			}
			maxLen = max(maxLen, curLen)
		}
	}
	return maxLen
}

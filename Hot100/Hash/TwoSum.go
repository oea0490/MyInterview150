package Hash

func TwoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	ans := make([]int, 2)
	for index, num := range nums {
		if _, ok := numToIndex[target-num]; ok {
			ans[0] = index
			ans[1] = numToIndex[target-num]
			return ans
		}
		numToIndex[num] = index
	}
	return ans
}

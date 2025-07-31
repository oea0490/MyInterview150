package Backtracking

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	subsetsDfs(nums, &ans, &path, 0)
	return ans
}

func subsetsDfs(nums []int, ans *[][]int, path *[]int, index int) {
	if index == len(nums) {
		*ans = append(*ans, append([]int{}, *path...))
		return
	}

	// 不选择当前元素
	subsetsDfs(nums, ans, path, index+1)
	// 选择当前元素
	*path = append(*path, nums[index])
	subsetsDfs(nums, ans, path, index+1)
	*path = (*path)[:len(*path)-1]
}

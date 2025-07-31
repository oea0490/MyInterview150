package Backtracking

func permute(nums []int) [][]int {
	visited := make([]bool, len(nums))
	ans := make([][]int, 0)
	path := make([]int, 0)
	permuteDfs(nums, &ans, &path, &visited, 0)
	return ans
}

func permuteDfs(nums []int, ans *[][]int, path *[]int, visited *[]bool, cur int) {
	if cur == len(nums) {
		*ans = append(*ans, append([]int{}, *path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if (*visited)[i] {
			continue
		}
		// 选择这一步
		(*visited)[i] = true
		*path = append(*path, nums[i])
		// 搜索下一步
		permuteDfs(nums, ans, path, visited, cur+1)
		// 恢复现场
		(*visited)[i] = false
		*path = (*path)[:len(*path)-1]
	}
}

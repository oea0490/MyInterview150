package Backtracking

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	path := make([]int, 0)
	combinationSumDfs(candidates, &ans, &path, target, 0)
	return ans
}

func combinationSumDfs(candidates []int, ans *[][]int, path *[]int, target int, index int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ans = append(*ans, tmp)
		return
	}

	for i := index; i < len(candidates); i++ {
		// 选择这一步
		*path = append(*path, candidates[i])
		// 搜索下一步
		combinationSumDfs(candidates, ans, path, target-candidates[i], i)
		// 恢复现场
		*path = (*path)[:len(*path)-1]
	}
}

package Backtracking

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]byte, 0)
	generateParenthesisDfs(0, 0, n, &ans, &path)
	return ans
}

func generateParenthesisDfs(left, right, n int, ans *[]string, path *[]byte) {
	if left == n && right == n {
		*ans = append(*ans, string(*path))
		return
	}
	if left < n {
		// 选择这一步
		*path = append(*path, '(')
		// 搜索下一步
		generateParenthesisDfs(left+1, right, n, ans, path)
		// 恢复现场
		*path = (*path)[:len(*path)-1]
	}
	if right < left {
		// 选择这一步
		*path = append(*path, ')')
		// 搜索下一步
		generateParenthesisDfs(left, right+1, n, ans, path)
		// 恢复现场
		*path = (*path)[:len(*path)-1]
	}
}

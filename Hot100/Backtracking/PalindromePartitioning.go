package Backtracking

func partition(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)

	// 1.DP确认子串是否是回文
	// 定义isPalindrome[i][j]为s[i:j]是否是回文串
	// 状态转移方程：isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i+1][j-1]
	// 初始状态：isPalindrome[i][i] = true
	isPalindrome := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		isPalindrome[i] = make([]bool, len(s))
		isPalindrome[i][i] = true
	}
	// 状态转移
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if i+1 <= j-1 {
				isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i+1][j-1]
			} else {
				isPalindrome[i][j] = s[i] == s[j]
			}
		}
	}

	// 2.分割回文串
	partitionDfs(s, &ans, &path, 0, isPalindrome)

	return ans
}

func partitionDfs(s string, ans *[][]string, path *[]string, start int, isPalindrome [][]bool) {
	if start >= len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*ans = append(*ans, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		if !isPalindrome[start][i] {
			continue
		}
		// 选择这一步
		*path = append(*path, s[start:i+1])
		// 搜索下一步
		partitionDfs(s, ans, path, i+1, isPalindrome)
		// 恢复现场
		*path = (*path)[:len(*path)-1]
	}
}

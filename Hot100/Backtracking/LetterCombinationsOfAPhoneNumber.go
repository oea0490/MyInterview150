package Backtracking

var phone = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var ans []string
	var path []byte
	letterCombinationsDfs(digits, &ans, &path, 0)
	return ans
}

func letterCombinationsDfs(digits string, ans *[]string, path *[]byte, index int) {
	if index == len(digits) {
		*ans = append(*ans, string(*path))
		return
	}
	for _, v := range phone[digits[index]] {
		// 选择这一步
		*path = append(*path, v)
		// 搜索下一步
		letterCombinationsDfs(digits, ans, path, index+1)
		// 恢复现场
		*path = (*path)[:len(*path)-1]
	}
}

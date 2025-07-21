package Substring

func MinWindow(s string, t string) string {
	minLen := 0x3f3f3f3f
	start := 0
	left, right := 0, 0
	need := make(map[byte]int)
	for _, c := range t {
		need[byte(c)]++
	}
	window := make(map[byte]int)
	for right < len(s) {
		// 入滑动窗口
		cur := s[right]
		window[cur]++
		for check(need, window) && left <= right {
			// 更新答案
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			// 出滑动窗口
			window[s[left]]--
			left++
		}
		right++
	}
	if minLen == 0x3f3f3f3f {
		return ""
	}
	return s[start : start+minLen]
}

func check(need, window map[byte]int) bool {
	for k, v := range need {
		if window[k] < v {
			return false
		}
	}
	return true
}

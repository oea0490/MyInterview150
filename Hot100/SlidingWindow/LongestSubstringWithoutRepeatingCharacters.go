package SlidingWindow

func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	left, right := 0, 0
	window := make(map[byte]bool)
	for right < len(s) {
		cur := s[right]
		for window[cur] {
			delete(window, s[left])
			left++
		}
		window[cur] = true
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}

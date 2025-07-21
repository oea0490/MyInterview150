package SlidingWindow

func FindAnagrams(s string, p string) []int {
	smap := make(map[byte]int)
	pmap := make(map[byte]int)
	for _, c := range p {
		pmap[byte(c)]++
	}
	var res []int
	left, right := 0, 0
	for right < len(s) {
		cur := s[right]
		smap[cur]++
		right++
		if right < len(p) {
			continue
		}
		if check(smap, pmap) {
			res = append(res, left)
		}
		smap[s[left]]--
		left++
	}
	return res
}

func check(smap, pmap map[byte]int) bool {
	for k, v := range pmap {
		if smap[k] != v {
			return false
		}
	}
	return true
}

package Hash

import "sort"

func GroupAnagrams(strs []string) [][]string {
	letterToAnagrams := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		letterToAnagrams[sortedStr] = append(letterToAnagrams[sortedStr], str)
	}
	var res [][]string
	for _, anagrams := range letterToAnagrams {
		res = append(res, anagrams)
	}
	return res
}

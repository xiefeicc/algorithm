package problem3

// abba
func lengthOfLongestSubstring(s string) int {
	var maxLen int
	var l, r int
	chars := make(map[byte]int)
	for ; r < len(s); r++ {
		c := s[r]
		if index, ok := chars[c]; ok && index >= l {
			maxLen = max(maxLen, r-l)
			l = index + 1
		}
		chars[c] = r
	}
	maxLen = max(maxLen, r-l)
	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

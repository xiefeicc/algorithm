package p5

// LongestPalindrome 最长回文
func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var palindrome string
	for i := 0; i < len(s)-1; i++ {
		l, r := i, i
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				break
			}
			r = j
		}

		for {
			if l-1 < 0 || r+1 >= len(s) {
				break
			}
			if s[l-1] != s[r+1] {
				break
			}
			l = l - 1
			r = r + 1
		}
		if r-l+1 > len(palindrome) {
			palindrome = s[l : r+1]
		}
	}
	return palindrome
}

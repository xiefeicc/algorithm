package p5

// LongestPalindrome 最长回文
func LongestPalindrome(s string) string {
	var palindrome string
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for ; r < len(s); r++ {
			if s[i] != s[r] {
				break
			}
		}
		r = r - 1

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

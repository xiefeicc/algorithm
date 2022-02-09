package p5

import "fmt"

// LongestPalindrome 最长回文
func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var palindrome string
	for i := 0; i < len(s)-1; i++ {
		// bbb
		if s[i] == s[i+1] {
			tmp := []byte{s[i]}
			for j := i + 1; j < len(s); j++ {
				if s[i] != s[j] {
					break
				}
				tmp = append(tmp, s[j])
			}
			palindrome = getMaxStr(palindrome, string(tmp))
			continue
		}
		// aba
		tmp := string(s[i])
		for n := 1; ; n++ {
			if i-n < 0 || i+n >= len(s) {
				break
			}
			if s[i-n] != s[i+n] {
				break
			}
			tmp = fmt.Sprintf("%s%s%s", string(s[i-n]), tmp, string(s[i+n]))
		}
		palindrome = getMaxStr(palindrome, tmp)
	}
	return palindrome
}

func getMaxStr(s1, s2 string) string {
	fmt.Printf("s1: %s, s2: %s, \n", s1, s2)
	if len(s2) > len(s1) {
		return s2
	}
	return s1
}

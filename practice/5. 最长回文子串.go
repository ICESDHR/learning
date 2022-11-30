package main

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := Palindrome(s, i, i)
		s2 := Palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func Palindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}

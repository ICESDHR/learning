package main

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		str1 := palindrome(s, i, i)
		str2 := palindrome(s, i, i+1)
		if len(str1) > len(res) {
			res = str1
		}
		if len(str2) > len(res) {
			res = str2
		}
	}
	return res
}

func palindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}

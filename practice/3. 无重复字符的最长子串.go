package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	slow, fast := 0, 1
	count := make(map[byte]int)
	count[s[0]]++
	res := 0
	for slow <= fast && fast < len(s) {
		if count[s[fast]] == 0 {
			count[s[fast]]++
			if fast-slow+1 > res {
				res = fast - slow + 1
			}
			fast++
			continue
		}

		for count[s[fast]] != 0 {
			count[s[slow]]--
			slow++
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

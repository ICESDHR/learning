package main

func lengthOfLongestSubstring(s string) int {
	result := 0
	dic := make(map[rune]int)
	start := 0
	for i, c := range []rune(s) {
		if index, ok := dic[c]; ok && index >= start {
			start = i + 1
		}
		dic[c] = i
		if i-start+1 > result {
			result = i - start + 1
		}
	}
	return result
}

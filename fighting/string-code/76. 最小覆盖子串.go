package main

func minWindow(s string, t string) string {
	memo := make(map[rune]int)
	for i, c := range t {
		memo[c] = i
	}
	res := ""
	left, right := 0, 0
	for right < len(s) {
		if _, ok := memo[s[right]]; !ok {
			right++

		}
	}
	return res
}

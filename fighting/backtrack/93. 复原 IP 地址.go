package main

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := []string{}
	track := []string{}
	backtrack(s, track, 0, &res)
	return res
}

func backtrack(s string, track []string, start int, res *[]string) {
	if len(track) == 4 && start == len(s) {
		tmp := make([]string, 4)
		copy(tmp, track)
		*res = append(*res, strings.Join(tmp, "."))
	}
	if len(track) == 4 && start < len(s) {
		return
	}
	for length := 1; length <= 3; length++ {
		if start+length-1 >= len(s) {
			return
		}
		if length != 1 && s[start] == '0' {
			return
		}
		str := s[start : start+length]
		if n, _ := strconv.Atoi(str); n > 255 {
			return
		}
		track = append(track, str)
		backtrack(s, track, start+length, res)
		track = track[:len(track)-1]
	}
}

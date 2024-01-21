package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}

	l := len(v1)
	for i := 0; i < l; i++ {
		num1, _ := strconv.Atoi(v1[i])
		num2, _ := strconv.Atoi(v2[i])
		if num1 == num2 {
			continue
		}
		if num1 < num2 {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

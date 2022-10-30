package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	res := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		tmp := 0
		if i >= 0 {
			tmp += int(num1[i] - '0')
		}
		if j >= 0 {
			tmp += int(num2[j] - '0')
		}
		res = strconv.Itoa((tmp+add)%10) + res
		add = (tmp + add) / 10
	}
	return res
}

func main() {
	fmt.Println(addStrings("", ""))
}

package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}
	res := ""
	add := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		tmp := 0
		if i >= 0 {
			tmp = int(num1[i] - '0')
		}
		if j >= 0 {
			tmp += int(num2[j] - '0')
		}
		tmp += add
		res = strconv.Itoa(tmp%10) + res
		add = tmp / 10
	}
	if add != 0 {
		res = strconv.Itoa(add) + res
	}
	return res
}

func main() {
	addStrings("11", "123")
}

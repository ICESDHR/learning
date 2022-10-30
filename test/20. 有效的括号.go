package main

import "fmt"

var tmp = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := tmp[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != tmp[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	if isValid("]") {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

package main

import "fmt"

func isValid(s string) bool {
	stack := []string{}
	tmp := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			stack = append(stack, string(s[i]))
			continue
		}
		if len(stack) == 0 || tmp[string(s[i])] != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid(")"))
}

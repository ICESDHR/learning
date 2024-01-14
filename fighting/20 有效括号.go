package main

func isValid(s string) bool {
	memo := make(map[byte]byte)
	memo['('] = ')'
	memo['['] = ']'
	memo['{'] = '}'

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			pop := stack[len(stack)-1]
			if memo[pop] != s[i] {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

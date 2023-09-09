package main

import (
	"fmt"
)

func main() {

	result := isValid("()[{}](}")

	fmt.Printf("result is %v\n", result)
}

func isValid(s string) bool {

	stack := make([]string, 0)

	for i := range s {
		switch str := string(s[i]); str {
		case "(":
			stack = append(stack, ")")
		case "[":
			stack = append(stack, "]")
		case "{":
			stack = append(stack, "}")
		default:
			if len(stack) == 0 || stack[len(stack)-1] != str {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

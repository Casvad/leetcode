package main

import "fmt"

func main() {

	result := generateParenthesis(3)

	fmt.Printf("result is: %v", result)

}

func generateParenthesis(n int) []string {
	if n == 0 {
		return make([]string, 0)
	}

	return recursiveGenerateParenthesis(n, 2, 1, "(")
}

func recursiveGenerateParenthesis(n, index, stackSize int, currentString string) []string {

	if index > n && stackSize == 0 {
		return []string{currentString}
	}

	openParenthesis := make([]string, 0)
	closedParenthesis := make([]string, 0)

	if index <= n {
		recursiveResult := recursiveGenerateParenthesis(n, index+1, stackSize+1, currentString+"(")
		openParenthesis = append(openParenthesis, recursiveResult...)
	}
	if stackSize > 0 {
		recursiveResult := recursiveGenerateParenthesis(n, index, stackSize-1, currentString+")")
		closedParenthesis = append(closedParenthesis, recursiveResult...)
	}

	return append(openParenthesis, closedParenthesis...)
}

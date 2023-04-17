package main

import (
	"fmt"
)

func main() {

	//runExample("babad")
	//runExample("cbbd")
	//runExample("bb")
	runExample("ccc")
}

func runExample(s string) {
	result := longestPalindrome(s)
	fmt.Printf("The result of %v is %v\n", s, result)
}

func longestPalindrome(s string) string {
	maxPalindrome := string(s[0])
	for i := 0; i < len(s); i++ {
		oddPalindrome := longestPalindromeWithIndex(&s, "odd", i)
		evenPalindrome := longestPalindromeWithIndex(&s, "even", i)
		if len(oddPalindrome) > len(maxPalindrome) {
			maxPalindrome = oddPalindrome
		}
		if len(evenPalindrome) > len(maxPalindrome) {
			maxPalindrome = evenPalindrome
		}
	}

	return maxPalindrome
}

func longestPalindromeWithIndex(s *string, mode string, i int) string {
	currentPalindrome := ""
	j := 1
	k := 0
	if mode == "odd" {
		k = 1
		currentPalindrome = string((*s)[i])
	}

	for true {

		if i+j > len(*s)-1 || i-k < 0 {
			break
		}

		if string((*s)[i+j]) == string((*s)[i-k]) {
			currentPalindrome = string((*s)[i-k]) + currentPalindrome + string((*s)[i+j])
		} else {
			break
		}
		j++
		k++
	}

	return currentPalindrome
}

package main

import (
	"fmt"
)

func main() {

	runExample("abcabcbb")
	runExample("bbbbb")
	runExample("pwwkew")
}

func runExample(s string) {
	result := lengthOfLongestSubstring(s)
	fmt.Printf("The result of %v is %d\n", s, result)
}

func lengthOfLongestSubstring(s string) int {
	longestSubstring := 0
	for i := 0; i < len(s); i++ {
		maxSubstring := getLongestSubstringForIndex(s, i)
		if maxSubstring > longestSubstring {
			longestSubstring = maxSubstring
		}
	}

	return longestSubstring
}

func getLongestSubstringForIndex(s string, index int) int {

	duplicatedMap := make(map[string]string)
	currentSubstring := 0
	for i := index; i < len(s); i++ {
		sString := string(s[i])
		_, found := duplicatedMap[sString]
		if !found {
			duplicatedMap[sString] = sString
			currentSubstring++
		} else {
			break
		}
	}
	return currentSubstring
}

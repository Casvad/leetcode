package main

import (
	"fmt"
)

func main() {
	runExample([]string{"flower", "flow", "flight"})
	runExample([]string{"dog", "racecar", "car"})
}

func runExample(strs []string) {
	result := longestCommonPrefix(strs)
	fmt.Printf("The result of %v is %v\n", strs, result)
}

func longestCommonPrefix(strs []string) string {
	result := ""
	index := 0
	if len(strs) == 0 {
		return ""
	}
	for {
		if len(strs[0]) <= index {
			break
		}
		toCompare := strs[0][index]
		canContinue := true
		for i := range strs {
			if len(strs[i]) <= index || strs[i][index] != toCompare {
				canContinue = false
				break
			}
		}
		if !canContinue {
			break
		}
		result += string(toCompare)
		index++
	}

	return result
}

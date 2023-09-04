package main

import (
	"fmt"
)

func main() {

	runExample("23")
	runExample("")
	runExample("2")
}

func runExample(input string) {
	result := letterCombinations(input)
	fmt.Printf("%v is %v\n", input, result)
}

func letterCombinations(digits string) []string {

	combinations := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	ans := make([]string, 0)

	for i := range digits {
		newAns := make([]string, 0)
		for _, j := range combinations[string(digits[i])] {

			if len(ans) == 0 {
				newAns = append(newAns, j)
			} else {
				for k := range ans {
					newAns = append(newAns, ans[k]+j)
				}
			}
		}

		ans = newAns
	}

	return ans
}

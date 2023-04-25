package main

import (
	"fmt"
)

func main() {

	runExample("PAYPALISHIRING", 3)
	runExample("PAYPALISHIRING", 4)
	runExample("A", 1)
	runExample("ABC", 1)
}

func runExample(s string, numRows int) {
	result := convert(s, numRows)
	fmt.Printf("The result of %v %v is %v\n", s, numRows, result)
}

func convert(s string, numRows int) string {
	if len(s) <= 2 || numRows == 1 {
		return s
	}

	stacks := make(map[int]string)

	i := 0
	operator := "inc"
	for _, subString := range s {
		if _, founded := stacks[i]; !founded {
			stacks[i] = ""
		}
		stacks[i] += string(subString)
		if operator == "inc" {
			i++
		} else {
			i--
		}
		if i >= numRows {
			i = i - 2
			operator = "dec"
		}
		if i < 0 {
			i = i + 2
			operator = "inc"
		}
	}

	res := ""
	for j := 0; j < numRows; j++ {
		if _, founded := stacks[j]; founded {
			res += stacks[j]
		}
	}
	return res
}

package main

import (
	"fmt"
)

func main() {

	runExample(121)
	runExample(-121)
	runExample(10)
}

func runExample(numRows int) {
	result := isPalindrome(numRows)
	fmt.Printf("The result of %v is %v \n", numRows, result)
}

func isPalindrome(x int) bool {

	digits := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9"}

	if x < 0 {
		return false
	}

	strResult := ""

	for x >= 10 {
		module := x % 10
		x = x / 10
		strResult = digits[module] + strResult
	}
	strResult = digits[x] + strResult

	i := 0
	j := len(strResult) - 1

	for i <= j {

		if strResult[i] != strResult[j] {
			return false
		}
		i++
		j--
	}

	return true
}

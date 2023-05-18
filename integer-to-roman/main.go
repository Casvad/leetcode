package main

import (
	"fmt"
)

func main() {

	//runExample(3)
	//runExample(58)
	//runExample(1994)
	runExample(499)
}

func runExample(numRows int) {
	result := intToRoman(numRows)
	fmt.Printf("The result of %v is %v\n", numRows, result)
}

func intToRoman(num int) string {
	romanNumbersIndex := []int{1000, 500, 100, 50, 10, 5, 1}
	romanNumbers := map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}

	stringResponse := ""
	ans := num
	repeatPossibility := 3

	for i := range romanNumbersIndex {
		k := romanNumbersIndex[i]
		v := romanNumbers[k]
		switch {
		case 900 <= ans && ans < 1000:
			ans = ans - 900
			stringResponse += "CM"
		case 400 <= ans && ans < 500:
			ans = ans - 400
			stringResponse += "CD"
		case 90 <= ans && ans < 100:
			ans = ans - 90
			stringResponse += "XC"
		case 40 <= ans && ans < 50:
			ans = ans - 40
			stringResponse += "XL"
		case ans == 9:
			ans = ans - 9
			stringResponse += "IX"
		case ans == 4:
			ans = ans - 4
			stringResponse += "IV"
		}
		if k <= ans {
			for j := 0; j < repeatPossibility; j++ {
				if ans >= k {
					ans = ans - k
					stringResponse += v
				}
			}
		}
	}

	return stringResponse
}

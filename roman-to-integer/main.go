package main

import (
	"fmt"
)

func main() {
	//runExample("III")
	//runExample("LVIII")
	runExample("MCMXCIV")
}

func runExample(numRows string) {
	result := romanToInt(numRows)
	fmt.Printf("The result of %v is %v\n", numRows, result)
}

func romanToInt(s string) int {
	romanNumbers := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	specialNumbers := map[string]int{
		"CM": 900,
		"CD": 400,
		"XC": 90,
		"XL": 40,
		"IX": 9,
		"IV": 4,
	}

	sum := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && specialNumbers[s[i:i+2]] != 0 {
			sum += specialNumbers[s[i:i+2]]
			i++
		} else {
			sum += romanNumbers[string(s[i])]
		}
	}

	return sum
}

package main

import (
	"fmt"
	"math"
)

func main() {

	runExample("20000000000000000000") //2147483647
	//runExample("")
	//runExample("     -42")
	//runExample("42")
	//runExample("4193 with words")
	//runExample("words and 987")
	//runExample("-91283472332") // -2147483648
}

func runExample(i string) {
	result := myAtoi(i)
	fmt.Printf("The result of %v is %v\n", i, result)
}

func myAtoi(s string) int {

	digits := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	i := 0
	for i < len(s) {
		if string(s[i]) != " " {
			break
		}
		i++
	}

	if i < len(s) && (string(s[i]) == "-" || string(s[i]) == "+") {
		i++
	}

	j := i
	numbersSum := 0
	upperLimit := 0
	if j > 0 && string(s[j-1]) == "-" {
		upperLimit = int(math.Pow(2, 31))
	} else {
		upperLimit = int(math.Pow(2, 31)) - 1
	}

	for i < len(s) {
		if _, ok := digits[string(s[i])]; !ok {
			break
		}

		numbersSum = 10*numbersSum + digits[string(s[i])]

		if numbersSum > upperLimit {
			numbersSum = upperLimit
			break
		}

		i++
	}

	if j > 0 && string(s[j-1]) == "-" {
		return -1 * numbersSum
	}

	return numbersSum
}

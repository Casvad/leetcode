package main

import (
	"fmt"
	"math"
)

func main() {

	runExample(1534236469)
	runExample(10)
	runExample(123)
	runExample(-123)
	runExample(120)
}

func runExample(i int) {
	result := reverse(i)
	fmt.Printf("The result of %v is %v\n", i, result)
}

func reverse(x int) int {
	sign := ""
	ans := 0
	numbersList := make([]int, 0)
	if x < 0 {
		sign = "-"
		x = -1 * x
	}

	for x >= 10 {
		toAdd := x % 10
		x = x / 10
		numbersList = append(numbersList, toAdd)
	}
	ans = x
	for i := len(numbersList) - 1; i >= 0; i-- {
		ans = numbersList[i]*int(math.Pow(10, float64(len(numbersList)-i))) + ans
	}

	if float64(ans) > math.Pow(2, 31) {
		return 0
	}

	if sign == "-" {
		ans = ans * -1
	}
	return ans
}

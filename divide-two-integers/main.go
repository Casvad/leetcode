package main

import (
	"fmt"
	"math"
)

func main() {
	dividend := -2147483648
	divisor := -1
	ans := divide(dividend, divisor)
	fmt.Printf("calculated=%d,real=%d\n", ans, dividend/divisor)
}

func divide(dividend int, divisor int) int {

	isPositive := dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0

	divisor = int(math.Abs(float64(divisor)))
	dividend = int(math.Abs(float64(dividend)))
	ans := 0
	d := dividend
	for d >= divisor {
		i := 0
		for d < divisor<<i {
			i++
		}
		ans += i + 1
		d -= divisor << i
	}

	if !isPositive {
		ans = -ans
	}

	return ans
}

package main

import (
	"fmt"
)

func main() {

	result := robbers([]int{50, 1, 1, 50})

	fmt.Printf("result is %v\n", result)
}

func max(number1, number2 int) int {
	if number1 > number2 {
		return number1
	}
	return number2
}

func robbers(houses []int) int {
	housesRobbed := make(map[int]int, 0)
	var dfs func(int) int
	dfs = func(index int) int {

		if index >= len(houses) {
			return 0
		}
		if amount, founded := housesRobbed[index]; founded {
			return amount
		}

		housesRobbed[index] = max(houses[index]+dfs(index+2), dfs(index+1))

		return housesRobbed[index]
	}

	return dfs(0)
}

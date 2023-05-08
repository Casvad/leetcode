package main

import (
	"fmt"
	"math"
)

func main() {

	runExample([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}

func runExample(height []int) {
	result := maxArea(height)
	fmt.Printf("The result of %v is %v\n", height, result)
}

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	maxCurArea := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))

	for i < j {
		if height[i] < height[j] {
			i++
		} else {
			j--
		}

		currentArea := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
		if currentArea > maxCurArea {
			fmt.Printf("Calculating with %v %v area -> %v \n", i, j, currentArea)
			maxCurArea = currentArea
		}
	}

	return maxCurArea
}

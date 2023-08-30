package main

import (
	"fmt"
)

func main() {

	//runExample([]int32{0, 0, 1, 5}, []int32{0, 1, 1, 0}) // [2 0 1 5]
	//runExample([]int32{0, 1, 1, 3, 3}, []int32{0, 1, 0, 0, 1}) // []
	//runExample([]int32{10}, []int32{0})
	runExample([]int32{3, 3, 3, 4, 4, 5, 6, 6, 7, 8}, []int32{1, 1, 0, 1, 0, 0, 0, 1, 0, 0})

}

func runExample(time []int32, direction []int32) {
	result := getTimes(time, direction)
	fmt.Printf("The result of %v is %v\n", time, result)
}

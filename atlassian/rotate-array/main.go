package main

import "fmt"

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6}, 2)
}

func rotate(arr []int, inputOrder int) {

	order := inputOrder
	if order >= len(arr) {
		order = inputOrder % len(arr)
	}

	pivot := len(arr) - order

	partialRotation(&arr, 0, pivot-1)
	partialRotation(&arr, pivot, len(arr)-1)
	partialRotation(&arr, 0, len(arr)-1)

	fmt.Printf("result is %v\n", arr)
}

func partialRotation(arr *[]int, start, end int) {

	for start < end {
		tmp := (*arr)[start]
		(*arr)[start] = (*arr)[end]
		(*arr)[end] = tmp
		start++
		end--
	}
}

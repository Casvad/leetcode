package main

import (
	"fmt"
)

func main() {

	result := QuickSort([]int{8, 2, 5, 3, 9, 4, 7, 6, 1})

	fmt.Printf("result is %v\n", result)
}

func QuickSort(arr []int) []int {

	quickSort(arr, 0, len(arr)-1)

	return arr
}

func quickSort(arr []int, start, end int) {

	if start < end {
		piv := pivot(arr, start, end)

		quickSort(arr, start, piv-1)
		quickSort(arr, piv+1, end)
	}
}

func pivot(arr []int, start, end int) int {

	piv := arr[end]
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < piv {
			i++
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
		}
	}
	i++
	arr[end] = arr[i]
	arr[i] = piv

	return i
}

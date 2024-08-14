package main

import "fmt"

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	quickSort(arr)
	fmt.Println("Sorted array:", arr)

	x := string(int(6 % 10))

	fmt.Println("Sorted array:", x)
}

func quickSort(arr []int) {

	if len(arr) <= 2 {
		return
	}

	left, right, pivot := 0, len(arr)-1, len(arr)/2
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[right], arr[left] = arr[left], arr[right]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

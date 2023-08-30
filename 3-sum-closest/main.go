package main

import (
	"fmt"
	"math"
)

func main() {

	runExample([]int{-1, 2, 1, -4}, 1)
}

func runExample(nums []int, target int) {
	result := threeSumClosest(nums, target)
	fmt.Printf("%v is %v\n", nums, result)
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func binarySearch(arr []int, number int) int {

	if len(arr) == 1 {
		return arr[0]
	}

	center := len(arr) / 2
	if arr[center] == number {
		return arr[center]
	}
	if arr[center] < number {
		return binarySearch(arr[center:], number)
	} else {
		return binarySearch(arr[:center], number)
	}
}

func threeSumClosest(nums []int, target int) int {
	nums2 := quickSortStart(nums)
	closest := -100
	calculated := false

	for i := 0; i < len(nums2)-2; i++ {
		for j := i + 1; j < len(nums2)-1; j++ {
			calculationClosest := binarySearch(nums2[j+1:], target-nums2[i]-nums2[j])

			if !calculated || math.Abs(float64(target-(nums2[i]+nums2[j]+calculationClosest))) < math.Abs(float64(target-closest)) {
				closest = nums2[i] + nums2[j] + calculationClosest
				calculated = true
			}
		}
	}

	return closest
}

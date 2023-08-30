package main

import "fmt"

func main() {

	runExample([]int{-1, 0, 1, 2, -1, -4})
}

func runExample(nums []int) {
	result := threeSum(nums)
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

func binarySearch(arr []int, number int) bool {

	if len(arr) == 0 {
		return false
	}

	center := len(arr) / 2
	if arr[center] == number {
		return true
	}
	if arr[center] < number {
		return binarySearch(arr[center+1:], number)
	} else {
		return binarySearch(arr[:center], number)
	}
}

func threeSum(nums []int) [][]int {
	nums2 := quickSortStart(nums)
	resultMap := make(map[string][]int, 0)
	result := make([][]int, 0)
	for i := 0; i < len(nums2)-2; i++ {
		for j := i + 1; j < len(nums2)-1; j++ {
			if binarySearch(nums2[j+1:], -nums2[i]-nums2[j]) {
				toAppend := []int{nums2[i], nums2[j], -nums2[i] - nums2[j]}
				resultMap[fmt.Sprintf("%d,%d,%d", toAppend[0], toAppend[1], toAppend[2])] = toAppend
			}
		}
	}

	for _, v := range resultMap {
		result = append(result, v)
	}

	return result
}

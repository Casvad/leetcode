package main

import "fmt"

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func isSortedDescending(nums []int, startIndex int) bool {

	for i := startIndex; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}

	return true
}

func nextPermutation(nums []int) {

	if len(nums) < 2 {
		return
	}

	pivot := 0
	pivotUsed := false

	for pivot <= len(nums)-2 {

		if nums[pivot] > nums[pivot+1] {

			switch {
			case isSortedDescending(nums, pivot):
				if pivot == 0 {
					quickSort(nums, 0, len(nums)-1)
				} else {
					pivotToExchange := len(nums) - 1
					for nums[pivotToExchange] <= nums[pivot-1] && pivotToExchange >= pivot {
						pivotToExchange--
					}
					nums[pivotToExchange], nums[pivot-1] = nums[pivot-1], nums[pivotToExchange]
					quickSort(nums, pivot, len(nums)-1)
				}
				pivotUsed = true
			default:

				for i := len(nums) - 1; i > pivot; i-- {
					if nums[i] > nums[i-1] {

						toSwap := i - 1
						for j := len(nums) - 1; j > toSwap; j-- {
							if nums[toSwap] < nums[j] {
								nums[toSwap], nums[j] = nums[j], nums[toSwap]
								break
							}
						}
						quickSort(nums, toSwap+1, len(nums)-1)
						pivotUsed = true
						pivotUsed = true
						break
					}
				}
			}
			break
		}
		pivot++
	}
	if !pivotUsed {
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				break
			}
		}
	}
}

func main() {
	ans := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", ans)
	for i := 1; i < 10; i++ {
		nextPermutation(ans)
		fmt.Printf("%v\n", ans)
	}
}

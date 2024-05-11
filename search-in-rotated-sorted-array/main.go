package main

import "fmt"

func binarySearch(init, end int, nums []int, target int) int {
	if init >= end {
		if nums[init] == target {
			return init
		}
		return -1
	}

	mid := (init + end) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return binarySearch(init, mid-1, nums, target)
	}

	return binarySearch(mid+1, end, nums, target)
}

func findKPoint(init, end int, nums []int) int {

	if init+1 >= end {

		if nums[init] > nums[end] {
			return init + 1
		}

		return 0
	}

	middle := (init + end) / 2
	if nums[middle] > nums[end] {
		return findKPoint(middle, end, nums)
	} else {
		return findKPoint(init, middle, nums)
	}
}

func search(nums []int, target int) int {

	k := findKPoint(0, len(nums)-1, nums)

	b1 := binarySearch(0, k, nums, target)
	if b1 != -1 {
		return b1
	}

	return binarySearch(k, len(nums)-1, nums, target)

}

func main() {

	testCases := [][]int{
		{4, 5, 6, 7, 0, 1, 2, 0, 4},
		{3, 1, 1, 1},
		{5, 1, 2, 3, 4, 4, 4},
	}

	for _, testCase := range testCases {
		result := search(testCase[:len(testCase)-2], testCase[len(testCase)-2])

		if result != testCase[len(testCase)-1] {
			fmt.Printf("Expected %v but was %v\n", testCase[len(testCase)-1], result)
		}
	}
}

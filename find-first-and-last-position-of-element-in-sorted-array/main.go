package main

import "fmt"

func binarySearch(init, end int, nums []int, target int, addMid int) int {
	if init >= end {
		if init < len(nums) && nums[init] == target {
			return init
		}
		return -1
	}

	mid := ((init + end) / 2) + addMid
	if nums[mid] == target {
		return mid
	}

	if nums[mid] > target {
		return binarySearch(init, mid-1, nums, target, addMid)
	}

	return binarySearch(mid+1, end, nums, target, addMid)
}

func findLimitPoint(init, end int, nums []int, asc bool, target int) int {

	if asc {
		if init == len(nums)-1 || nums[init+1] != target {
			return init
		}

		bs := binarySearch(init, len(nums)-1, nums, target, 1)

		return findLimitPoint(bs, len(nums)-1, nums, asc, target)
	} else {

		if end <= 0 || nums[end-1] != target {
			return end
		}

		bs := binarySearch(0, end, nums, target, 0)

		return findLimitPoint(0, bs, nums, asc, target)
	}
}

func searchRange(nums []int, target int) []int {

	k := binarySearch(0, len(nums)-1, nums, target, 0)

	if k == -1 {
		return []int{-1, -1}
	}

	return []int{
		findLimitPoint(0, k, nums, false, target),
		findLimitPoint(k, len(nums)-1, nums, true, target),
	}
}

func main() {

	//fmt.Printf("Result: %v", searchRange([]int{}, 0))
	//fmt.Printf("Result: %v\n", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Printf("Result: %v\n", searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3))
}

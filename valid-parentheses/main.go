package main

import (
	"fmt"
	"sort"
)

func main() {

	//result := fourSum([]int{1, 0, -1, 0 ,-2, 2}, 0)
	//result := fourSum([]int{2,2,2,2,2}, 8)
	result := fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)
	//[-2,-1,1,2]
	//[[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

	fmt.Printf("result is %v\n", result)
}

func binarySearch(nums []int, start, end, toFind int) bool {
	if start > end {
		return false
	}
	middle := (end + start) / 2
	if nums[middle] == toFind {
		return true
	}
	if nums[middle] > toFind {
		return binarySearch(nums, start, middle-1, toFind)
	}

	return binarySearch(nums, middle+1, end, toFind)
}

func fourSum(nums []int, target int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := make([][]int, 0)
	optimizationMap := make(map[string]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				numberToFind := target - nums[i] - nums[k] - nums[j]
				mapKey := fmt.Sprintf("%d,%d,%d,%d", nums[i], nums[j], nums[k], numberToFind)
				if _, found := optimizationMap[mapKey]; !found {
					optimizationMap[mapKey] = 1
					if binarySearch(nums, k+1, len(nums)-1, numberToFind) {
						ans = append(ans, []int{nums[i], nums[j], nums[k], numberToFind})
					}
				}

			}
		}
	}

	return ans
}

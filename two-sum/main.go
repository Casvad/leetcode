package main

import "fmt"

func main() {

	//runExample([]int{2, 7, 11, 15}, 9)
	runExample([]int{-1, -2, -3, -4, -5}, -8)
}

func runExample(nums []int, target int) {
	result := twoSum(nums, target)
	fmt.Printf("%v is %v\n", nums, result)
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

//This solution only is valid for sorted nums and positive numbers
func twoSumV1(nums []int, target int) []int {

	startIndex := 0
	endIndex := len(nums) - 1

	for {
		currentSum := nums[startIndex] + nums[endIndex]
		if currentSum == target {
			return []int{startIndex, endIndex}
		}
		if currentSum < target {
			startIndex++
		} else {
			endIndex--
		}
	}
}

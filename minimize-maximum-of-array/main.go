package main

import (
	"fmt"
)

func main() {

	//runExample([]int{3, 7, 1, 6})                                                                                                                                                                                        //5
	//runExample([]int{10, 1})                                                                                                                                                                                             //10
	//runExample([]int{13, 13, 20, 0, 8, 9, 9})                                                                                                                                                                            //16
	//runExample([]int{6, 9, 3, 8, 14})                                                                                                                                                                                    //8
	//runExample([]int{4, 7, 2, 2, 9, 19, 16, 0, 3, 15})                                                                                                                                                                   //9
	runExample([]int{439, 228, 482, 150, 231, 209, 991, 125, 453, 407, 670, 491, 300, 125, 285, 749, 350, 411, 527, 768}) //439                                                                            //9
	//runExample([]int{621, 277, 638, 960, 972, 544, 994, 865, 352, 801, 171, 973, 959, 569, 182, 367, 471, 789, 929, 707, 735, 511, 781, 146, 855, 947, 82, 29, 671, 28, 570, 154, 651, 168, 28, 27, 404, 811, 649, 228}) //734
}

func runExample(input []int) {
	result := minimizeArrayValue(input)
	fmt.Printf("INPUT => %v result %v\n", input, result)
}

func minimizeArrayValue(nums []int) int {
	maxElementIndex := findMaxElementIndex(nums)
	maxValue := nums[maxElementIndex]
	minElementIndex := findMinimumIndexBefore(nums, maxElementIndex)

	for {
		cof := (nums[maxElementIndex] - nums[minElementIndex]) / 2

		fmt.Printf("%v taking %d and %d with cof %d \n", nums, nums[maxElementIndex], nums[minElementIndex], cof)
		nums[maxElementIndex] = nums[maxElementIndex] - cof
		nums[minElementIndex] = nums[minElementIndex] + cof

		newMaxValIndex := findMaxElementIndex(nums)
		newMinValIndex := findMinimumIndexBefore(nums, maxElementIndex)
		if newMaxValIndex == maxElementIndex && newMinValIndex == minElementIndex {
			break
		}
		maxElementIndex = newMaxValIndex
		maxValue = nums[maxElementIndex]
		minElementIndex = newMinValIndex
	}

	return maxValue
}

func findMaxElementIndex(nums []int) int {
	maxElement := nums[0]
	maxElementIndex := 0

	for i := 1; i < len(nums); i++ {
		if maxElement < nums[i] {
			maxElement = nums[i]
			maxElementIndex = i
		}
	}

	return maxElementIndex
}

func findMinimumIndexBefore(nums []int, limit int) int {
	minElement := nums[0]
	minElementIndex := 0

	for i := 0; i < limit; i++ {
		if minElement > nums[i] {
			minElement = nums[i]
			minElementIndex = i
		}
	}

	return minElementIndex
}

///This solution has speed problems because of the massive iterations to reach for the solution
func minimizeArrayValueV1(nums []int) int {

	for {
		indexToDecrease := 1
		for i := 2; i < len(nums); i++ {
			if nums[i] > nums[indexToDecrease] && nums[i-1] < nums[i] {
				indexToDecrease = i
			}
		}
		if nums[indexToDecrease] <= nums[indexToDecrease-1] {
			break
		}
		nums[indexToDecrease] = nums[indexToDecrease] - 1
		nums[indexToDecrease-1] = nums[indexToDecrease-1] + 1
	}

	maxNumber := -1

	for _, num := range nums {
		if num > maxNumber {
			maxNumber = num
		}
	}

	return maxNumber
}

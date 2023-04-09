package main

import (
	"fmt"
)

func main() {

	//runExample([]int{1, 3}, []int{2}) 2
	//runExample([]int{1, 2}, []int{3, 4}) 2.5
	//runExample([]int{1, 3}, []int{2, 7}) 3.5
	//runExample([]int{2}, []int{1, 3, 4, 5}) //3
	runExample([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4}, []int{1, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}) //3

}

func runExample(nums1 []int, nums2 []int) {
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("The result of %v is %f\n", nums1, result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	sortedArray := sortArrays(nums1, nums2)
	if len(sortedArray)%2 == 0 && len(sortedArray) > 1 {
		return (float64(sortedArray[len(sortedArray)/2]) + float64(sortedArray[(len(sortedArray)/2)-1])) / 2
	} else {
		return float64(sortedArray[len(sortedArray)/2])
	}
}

func sortArrays(nums1 []int, nums2 []int) []int {

	if len(nums2) == 0 {
		return nums1
	}
	if len(nums1) == 0 {
		return nums2
	}
	if nums1[0] >= nums2[len(nums2)-1] {
		return joinArray(nums2, nums1)
	}
	if nums1[len(nums1)-1] <= nums2[0] {
		return joinArray(nums1, nums2)
	}

	nums1Init := nums1[0 : len(nums1)/2]
	nums1End := nums1[len(nums1)/2:]
	nums2Init := nums2[0 : len(nums2)/2]
	nums2End := nums2[len(nums2)/2:]

	sortedInit := sortArrays(nums1Init, nums2Init)
	sortedEnd := sortArrays(nums1End, nums2End)

	if sortedInit[len(sortedInit)-1] <= sortedEnd[0] {
		return joinArray(sortedInit, sortedEnd)
	}

	solveArray := make([]int, 0)
	i := 0
	j := 0

	for i < len(sortedInit) || j < len(sortedEnd) {
		if i >= len(sortedInit) {
			solveArray = append(solveArray, sortedEnd[j])
			j++
		} else if j >= len(sortedEnd) {
			solveArray = append(solveArray, sortedInit[i])
			i++
		} else if sortedInit[i] < sortedEnd[j] {
			solveArray = append(solveArray, sortedInit[i])
			i++
		} else {
			solveArray = append(solveArray, sortedEnd[j])
			j++
		}
	}

	return solveArray
}

func joinArray(nums1 []int, nums2 []int) []int {

	r := append([]int{}, nums1...)

	return append(r, nums2...)
}

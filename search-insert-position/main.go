package main

func searchInsert(nums []int, target int) int {

	return binarySearch(0, len(nums)-1, target, nums)
}

func binarySearch(startIndex, endIndex, target int, nums []int) int {
	middle := (startIndex + endIndex) / 2

	if target == nums[middle] {
		return middle
	}

	if startIndex >= endIndex {
		if nums[startIndex] > target {
			return startIndex
		}
		return startIndex + 1
	}

	if nums[middle] < target {

		return binarySearch(middle+1, endIndex, target, nums)
	}

	return binarySearch(startIndex, middle-1, target, nums)

}

func main() {
	x := searchInsert([]int{1, 3, 5, 6}, 2)

	println(x)
}

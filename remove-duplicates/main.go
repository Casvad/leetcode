package main

func removeDuplicates(nums []int) int {

	mapCheck := make(map[int]int, 0)
	k := 0
	for _, i := range nums {
		if _, found := mapCheck[i]; !found {
			mapCheck[i] = i
			nums[k] = i
			k++
		}
	}

	return k
}

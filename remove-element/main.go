package main

func main() {

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	x := removeElement(nums, 2)

	println(x)
}

func removeElement(nums []int, val int) int {
	numsSize := len(nums)
	for i := 0; i < numsSize; i++ {
		num := nums[i]
		if num == val {
			for j := i; j < numsSize-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			numsSize--
			i--
		}
	}

	return numsSize
}

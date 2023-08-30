package main

import (
	"fmt"
)

func getTimes(time []int32, direction []int32) []int32 {
	timesResponse := make([]int32, len(time))
	inQueueIndex := -1
	outQueueIndex := -1
	lastUsage := int32(-2)
	lastDirection := 0
	iterationTime := int32(0)

	outLogic := func() {
		if iterationTime < time[outQueueIndex] {
			iterationTime = time[outQueueIndex]
		}
		timesResponse[outQueueIndex] = iterationTime
		lastDirection = 1
		lastUsage = iterationTime
		for i := outQueueIndex + 1; i <= len(time); i++ {
			if i == len(time) {
				outQueueIndex = len(time)
			} else {
				if direction[i] == 1 {
					outQueueIndex = i
					break
				}
			}
		}
		iterationTime++
	}

	inLogic := func() {
		if iterationTime < time[inQueueIndex] {
			iterationTime = time[inQueueIndex]
		}
		timesResponse[inQueueIndex] = iterationTime
		lastDirection = 0
		lastUsage = iterationTime
		for i := inQueueIndex + 1; i <= len(time); i++ {
			if i == len(time) {
				inQueueIndex = len(time)
			} else {
				if direction[i] == 0 {
					inQueueIndex = i
					break
				}
			}
		}
		iterationTime++
	}

	for i := 0; i < len(time); i++ {
		if inQueueIndex > -1 && outQueueIndex > -1 {
			break
		}
		if inQueueIndex == -1 {
			if direction[i] == 0 {
				inQueueIndex = i
			}
		}
		if outQueueIndex == -1 {
			if direction[i] == 1 {
				outQueueIndex = i
			}
		}
	}
	if outQueueIndex == -1 {
		outQueueIndex = len(time)
	}
	if inQueueIndex == -1 {
		inQueueIndex = len(time)
	}

	for {
		if inQueueIndex >= len(time) && outQueueIndex >= len(time) {
			return timesResponse
		}
		if inQueueIndex >= len(time) {
			for i := outQueueIndex; i < len(time); i++ {
				if direction[i] == 1 {
					if iterationTime < time[i] {
						iterationTime = time[i]
					}
					timesResponse[i] = iterationTime
					iterationTime++
				}
			}
			break
		}
		if outQueueIndex >= len(time) {
			for i := inQueueIndex; i < len(time); i++ {
				if direction[i] == 0 {
					if iterationTime < time[i] {
						iterationTime = time[i]
					}
					timesResponse[i] = iterationTime
					iterationTime++
				}
			}
			break
		}
		if time[outQueueIndex] == time[inQueueIndex] {
			if outQueueIndex < len(time) && (inQueueIndex >= len(time) || lastUsage < iterationTime-1 || lastDirection == 1) {
				outLogic()
				if inQueueIndex < len(time) {
					time[inQueueIndex]++
				}
			} else {
				inLogic()
				if outQueueIndex < len(time) {
					time[outQueueIndex]++
				}
			}
		} else {
			if time[outQueueIndex] < time[inQueueIndex] {
				outLogic()
			} else {
				inLogic()
			}
		}
	}

	return timesResponse
}

func main() {

	//runExample([]int32{0, 0, 1, 5}, []int32{0, 1, 1, 0}) // [2 0 1 5]
	//runExample([]int32{0, 1, 1, 3, 3}, []int32{0, 1, 0, 0, 1}) // []
	//runExample([]int32{10}, []int32{0})
	runExample([]int32{3, 3, 3, 4, 4, 5, 6, 6, 7, 8}, []int32{1, 1, 0, 1, 0, 0, 0, 1, 0, 0})

}

func runExample(time []int32, direction []int32) {
	result := getTimes(time, direction)
	fmt.Printf("The result of %v is %v\n", time, result)
}

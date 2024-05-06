package main

import (
	"fmt"
	"math"
)

func trapV1(height []int) int {

	startIndex := 0
	filledLevel := make([]int, len(height))
	for startIndex < len(height)-1 {
		if height[startIndex] < 1 {
			startIndex++
			continue
		}
		endIndex := startIndex + 1

		for endIndex <= len(height)-1 {
			if height[endIndex] < 1 {
				endIndex++
				continue
			}

			waterLevel := int(math.Min(float64(height[startIndex]), float64(height[endIndex])))
			for i := startIndex + 1; i < endIndex; i++ {
				newFillLevel := waterLevel - height[i]
				filledLevel[i] = int(math.Max(float64(filledLevel[i]), float64(newFillLevel)))
			}

			if height[startIndex] < height[endIndex] {
				startIndex = endIndex - 1
				break
			}
			endIndex += 1
		}

		startIndex += 1
	}
	trappedWater := 0
	for i := 0; i < len(filledLevel); i++ {
		trappedWater += filledLevel[i]
	}
	return trappedWater
}

func trapV0(height []int) int {

	startIndex := 0
	filledLevel := 0
	filledHeight := make([]int, len(height))
	for startIndex < len(height)-1 {
		if height[startIndex] < 1 {
			startIndex++
			continue
		}
		previousEndIndex := startIndex
		endIndex := startIndex + 1

		for endIndex <= len(height)-1 {
			if height[endIndex] < 1 {
				endIndex++
				continue
			}

			waterLevel := int(math.Min(float64(height[startIndex]), float64(height[endIndex])))
			for i := previousEndIndex + 1; i < endIndex; i++ {
				if waterLevel >= height[i] && filledHeight[i] < waterLevel {
					filledLevel += waterLevel - height[i]
					filledHeight[i] = waterLevel
				}
			}
			if waterLevel-height[previousEndIndex] > 0 && previousEndIndex < endIndex {
				filledLevel += (previousEndIndex - startIndex) * (waterLevel - height[previousEndIndex])
			}

			if height[startIndex] <= height[endIndex] {
				startIndex = endIndex - 1
				break
			} else {
				if previousEndIndex == startIndex || height[previousEndIndex] <= height[endIndex] {
					previousEndIndex = endIndex
				}
			}
			endIndex += 1
		}

		startIndex = endIndex
	}

	return filledLevel
}

func trapOK(height []int) int {

	startIndex := 0
	filledLevel := 0
	for startIndex < len(height)-1 {
		if height[startIndex] < 1 {
			startIndex++
			continue
		}
		for startIndex < len(height)-2 {
			if height[startIndex] >= height[startIndex+1] {
				break
			}
			startIndex++
		}
		currentLevelWater := 0
		endIndex := startIndex + 2

		for endIndex <= len(height)-1 {
			if height[endIndex] < 1 || height[endIndex] < height[endIndex-1] {
				endIndex++
				continue
			}

			currentLevelWater = int(math.Max(
				float64(currentLevelWater),
				float64(int(math.Min(float64(height[startIndex]), float64(height[endIndex])))),
			))

			if height[startIndex] <= height[endIndex] {
				break
			}

			shouldLimitPool := true
			for i := endIndex + 1; i <= len(height)-1; i++ {
				if height[i] >= height[endIndex] {
					shouldLimitPool = false
					break
				}
			}
			if shouldLimitPool {
				break
			}

			endIndex++
		}

		if endIndex >= len(height) {
			endIndex = len(height) - 1
		}

		for i := startIndex + 1; i < endIndex; i++ {
			if currentLevelWater-height[i] > 0 {
				filledLevel += currentLevelWater - height[i]
			}
		}

		startIndex = endIndex
	}

	return filledLevel
}

func trap(height []int) int {

	startIndex := 0
	filledLevel := 0
	for startIndex < len(height)-2 {
		if height[startIndex] < 1 || height[startIndex] < height[startIndex+1] {
			startIndex++
			continue
		}
		currentLevelWater := 0
		endIndex := startIndex + 2

		for endIndex <= len(height)-1 {
			if height[endIndex] < 1 || height[endIndex] < height[endIndex-1] {
				endIndex++
				continue
			}

			currentLevelWater = int(math.Max(
				float64(currentLevelWater),
				float64(int(math.Min(float64(height[startIndex]), float64(height[endIndex])))),
			))

			if height[startIndex] <= height[endIndex] {
				break
			}

			shouldLimitPool := true
			for i := endIndex + 1; i <= len(height)-1; i++ {
				if height[i] >= height[endIndex] {
					shouldLimitPool = false
					break
				}
			}
			if shouldLimitPool {
				break
			}

			endIndex++
		}

		for i := startIndex + 1; i < endIndex; i++ {
			if currentLevelWater-height[i] > 0 {
				filledLevel += currentLevelWater - height[i]
			}
		}

		startIndex = endIndex
	}

	return filledLevel
}

func main() {

	testsCases := [][]int{
		{4, 2, 3, 0, 1},
		{1000, 999, 998, 997, 996, 995, 994, 993, 992, 991, 990, 989, 988, 987, 986, 985, 984, 983, 982, 981, 980, 979, 978, 977, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966, 966},
		{0, 1, 2, 0, 3, 0, 1},
		{4, 4, 4, 7, 1, 0},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{2, 1, 0, 2},
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{5, 4, 1, 2},
		{6, 4, 2, 0, 3, 2, 0, 3, 5},
	}

	results := []int{
		2, 0, 3, 0, 6, 9, 3, 6, 1, 21,
	}

	for i := range testsCases {

		res := trap(testsCases[i])

		if res != results[i] {
			fmt.Printf("Error on case %v expecting %v but was %v\n", testsCases[i], results[i], res)
		}
	}
}

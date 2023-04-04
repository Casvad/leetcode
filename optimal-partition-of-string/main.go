package main

import (
	"fmt"
	"strings"
)

func main() {

	runExample("lvkmzlaeaxbprczpfarnlaptfvmutkfsatyywnxpmkpduwoqeeiltbdjipwihqi")
	runExample("abacaba")
	runExample("ssssss")
	runExample("a")
}

func runExample(input string) {
	result := partitionString(input)
	fmt.Printf("The partition string of %s is %d\n", input, result)
}

///this not uses brute force because is not necessary to check all the possible solutions
///checking the len of s in O(n) time is valid because is the minimum possible string (giving priority to reduce sub strings lists)
func partitionString(s string) int {

	sMap := make(map[string]string)
	res := 0

	for i := range s {
		if _, ok := sMap[string(s[i])]; ok {
			res++
			sMap = map[string]string{string(s[i]): string(s[i])}
		} else {
			sMap[string(s[i])] = string(s[i])
		}
	}

	if len(sMap) > 0 {
		res++
	}

	return res
}

///This example is v1 brute force checking all possible solutions
func partitionStringV1(s string) int {
	subStrings := partitionStringRecursive(s, string(s[0]), 1, []int{1})

	min := subStrings[0]
	for _, subStringValue := range subStrings {
		if subStringValue < min {
			min = subStringValue
		}
	}

	return min
}

func partitionStringRecursive(s, subString string, index int, currentSubStrings []int) []int {

	if index >= len(s) {
		return currentSubStrings[:]
	} else {

		if strings.Contains(subString, string(s[index])) {
			return partitionStringRecursive(s, string(s[index]), index+1, increaseSlice(currentSubStrings))
		} else {
			withElementResult := partitionStringRecursive(s, subString+string(s[index]), index+1, currentSubStrings)
			withoutElementResult := partitionStringRecursive(s, string(s[index]), index+1, increaseSlice(currentSubStrings))

			withElementResult = append(withElementResult, withoutElementResult...)

			return withElementResult[:]
		}
	}
}

func increaseSlice(sliceInt []int) []int {
	newSlice := make([]int, len(sliceInt))
	for i, v := range sliceInt {
		newSlice[i] = v + 1
	}

	return newSlice
}

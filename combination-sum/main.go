package main

import "fmt"

func getKey(index, currentSum int, currentCandidates []int) string {

	return fmt.Sprintf("%v_%v_%v", index, currentSum, currentCandidates)
}

func getPossibleCandidate(candidates []int, target, index, currentSum int, currentCandidates []int, persistence map[string]bool) [][]int {

	ans := make([][]int, 0)
	key := getKey(index, currentSum, currentCandidates)

	if index >= len(candidates) || currentSum > target || persistence[key] {
		return ans
	}

	persistence[key] = true

	if target == currentSum+candidates[index] {
		ans = append(ans, append(append([]int{}, currentCandidates...), candidates[index]))
	}

	withoutCounting := getPossibleCandidate(candidates, target, index+1, currentSum, append([]int{}, currentCandidates...), persistence)
	countingNotAdd := getPossibleCandidate(candidates, target, index, currentSum+candidates[index], append(append([]int{}, currentCandidates...), candidates[index]), persistence)
	countingAdd := getPossibleCandidate(candidates, target, index+1, currentSum+candidates[index], append(append([]int{}, currentCandidates...), candidates[index]), persistence)

	return append(append(append(ans, countingNotAdd...), countingAdd...), withoutCounting...)
}

func combinationSum(candidates []int, target int) [][]int {
	return getPossibleCandidate(candidates, target, 0, 0, make([]int, 0), make(map[string]bool))
}

func main() {

	//fmt.Printf("result: %v", combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Printf("result: %v\n", combinationSum([]int{2, 5, 8, 4}, 10))
	//fmt.Printf("result: %v", combinationSum([]int{2, 10, 5, 4}, 10))
}

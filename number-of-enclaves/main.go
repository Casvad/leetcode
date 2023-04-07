package main

import (
	"fmt"
)

func main() {

	//runExample([][]int{
	//	{0, 0, 0, 0},
	//	{1, 0, 1, 0},
	//	{0, 1, 1, 0},
	//	{0, 0, 0, 0},
	//}) //3

	//runExample([][]int{
	//	{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
	//	{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
	//	{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
	//	{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
	//	{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
	//	{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
	//	{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
	//	{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
	//	{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}) //3

	runExample([][]int{{0}, {1}, {1}, {0}})
}

func runExample(input [][]int) {
	result := numEnclaves(input)
	fmt.Printf("The result of %v is %d\n", input, result)
}

func numEnclaves(grid [][]int) int {
	//top row
	i := 0
	for j := 0; j < len(grid[i]); j++ {
		checkIslandsNear(i, j, grid)
	}
	//bottom row
	i = len(grid) - 1
	for j := 0; j < len(grid[i]); j++ {
		checkIslandsNear(i, j, grid)
	}
	//left row
	for i = 0; i < len(grid); i++ {
		checkIslandsNear(i, 0, grid)
	}
	//right row
	for i = 0; i < len(grid); i++ {
		checkIslandsNear(i, len(grid[i])-1, grid)
	}

	return countNotReachIsland(grid)
}

func countNotReachIsland(grid [][]int) int {

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
			}
		}
	}

	return count
}

func checkIslandsNear(i, j int, grid [][]int) {

	if grid[i][j] == 1 {
		grid[i][j] = 2

		//up
		if i-1 > 0 {
			checkIslandsNear(i-1, j, grid)
		}
		//down
		if i+1 < len(grid) {
			checkIslandsNear(i+1, j, grid)
		}
		//left
		if j-1 > 0 {
			checkIslandsNear(i, j-1, grid)
		}
		//right
		if j+1 < len(grid[i]) {
			checkIslandsNear(i, j+1, grid)
		}
	}
}

func printPrettyMatrix(grid [][]int) {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf(" %v ", grid[i][j])
		}
		fmt.Printf("\n")
	}
}

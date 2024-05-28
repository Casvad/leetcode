package main

import "encoding/json"

type matrix [][]string

func main() {

	x := `[
    [".",".",".",".","5",".",".","1","."],
    [".","4",".","3",".",".",".",".","."],
    [".",".",".",".",".","3",".",".","1"],
    ["8",".",".",".",".",".",".","2","."],
    [".",".","2",".","7",".",".",".","."],
    [".","1","5",".",".",".",".",".","."],
    [".",".",".",".",".","2",".",".","."],
    [".","2",".","9",".",".",".",".","."],
    [".",".","4",".",".",".",".",".","."]
]`
	var req matrix

	json.Unmarshal([]byte(x), &req)

	println(isValidSudoku(toByteArray(req)))
}

func toByteArray(input [][]string) [][]byte {

	output := make([][]byte, len(input))

	for i := range output {
		output[i] = make([]byte, len(input[i]))
		for j := range output[i] {
			output[i][j] = input[i][j][0]
		}
	}

	return output
}

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		if !isValidRow(i, board) || !isValidColumn(i, board) {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isValidSquare(i*3, j*3, board) {
				return false
			}
		}
	}

	return true
}

func isValidRow(row int, board [][]byte) bool {

	validMap := make(map[string]bool)
	for i := range board {
		if _, found := validMap[string(board[row][i])]; found {
			return false
		}

		if string(board[row][i]) != "." {
			validMap[string(board[row][i])] = true
		}

	}

	return true
}

func isValidColumn(column int, board [][]byte) bool {

	validMap := make(map[string]bool)
	for i := range board {
		if _, found := validMap[string(board[i][column])]; found {
			return false
		}
		if string(board[i][column]) != "." {
			validMap[string(board[i][column])] = true
		}

	}

	return true
}

func isValidSquare(rowStart, columnStart int, board [][]byte) bool {

	validMap := make(map[string]bool)
	for i := rowStart; i < rowStart+3; i++ {
		for j := columnStart; j < columnStart+3; j++ {
			if _, found := validMap[string(board[i][j])]; found {
				return false
			}
			if string(board[i][j]) != "." {
				validMap[string(board[i][j])] = true
			}
		}
	}

	return true
}

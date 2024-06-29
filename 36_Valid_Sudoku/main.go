package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))

}
func isValidSudoku(board [][]byte) bool {
	validSectors := make([]map[byte]bool, 9)
	validColumns := make([]map[byte]bool, 9)
	validRows := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		validSectors[i] = make(map[byte]bool, 9)
		validColumns[i] = make(map[byte]bool, 9)
		validRows[i] = make(map[byte]bool, 9)
	}
	s := 0
	for r := 0; r < len(board); r++ {
		s = (r / 3) * 3
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == '.' {
				continue
			}
			if validSectors[s+c/3][board[r][c]] {
				return false
			}
			validSectors[s+c/3][board[r][c]] = true
			if validRows[c][board[r][c]] {
				return false
			}
			validRows[c][board[r][c]] = true

			if validColumns[r][board[r][c]] {
				return false
			}
			validColumns[r][board[r][c]] = true
		}
	}
	return true
}

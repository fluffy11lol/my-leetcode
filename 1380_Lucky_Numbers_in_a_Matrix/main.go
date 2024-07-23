package main

import "math"

func main() {

}
func luckyNumbers(matrix [][]int) []int {
	luckyNums := make([]int, 0)
	type place struct {
		i, j int
	}
	minInRows := make(map[place]bool, len(matrix[0]))
	maxInCols := make(map[place]bool, len(matrix))
	curMax, curMin := 0, 0
	var curPlace place
	for i := range matrix {
		curMin = math.MaxInt
		for j := range matrix[0] {
			if matrix[i][j] < curMin {
				curMin = matrix[i][j]
				curPlace = place{i, j}
			}
		}
		minInRows[curPlace] = true
	}
	for j := range matrix[0] {
		curMax = 0
		for i := range matrix {
			if matrix[i][j] > curMax {
				curMax = matrix[i][j]
				curPlace = place{i, j}
			}
		}
		maxInCols[curPlace] = true
		if minInRows[curPlace] {
			luckyNums = append(luckyNums, matrix[curPlace.i][j])
		}
	}
	return luckyNums
}

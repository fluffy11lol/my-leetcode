package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
}
func countNegatives(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	sumOfNeg := 0
	jCrit := m
	for i := 0; i < n; i++ {
		j := 0
		for ; j < jCrit && grid[i][j] >= 0; j++ {
		}
		sumOfNeg += m - j
		jCrit = j
	}
	return sumOfNeg
}

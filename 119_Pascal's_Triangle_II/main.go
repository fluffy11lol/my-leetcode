package main

import "fmt"

func main() {
	fmt.Println(getRow(5))

}
func getRow(rowIndex int) []int {
	prevRow := []int{1}
	curRow := []int{1, 1}
	if rowIndex == 0 {
		return prevRow
	}
	if rowIndex == 1 {
		return curRow
	}
	i := 2
	for i <= rowIndex {
		prevRow = curRow
		curRow = make([]int, i/2+1+i&1)
		curRow[0] = 1
		for j := 1; j <= i/2; j++ {
			curRow[j] = prevRow[j-1] + prevRow[j]
		}
		if i&1 == 1 {
			curRow[len(curRow)-1] = curRow[len(curRow)-2]
		}
		i++
	}

	resRow := make([]int, rowIndex+1)
	copy(resRow, curRow)
	for i := rowIndex/2 + 1; i <= rowIndex; i++ {
		resRow[i] = resRow[len(resRow)-i-1]
	}

	return resRow
}

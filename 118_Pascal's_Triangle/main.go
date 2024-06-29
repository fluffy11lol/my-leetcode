package main

import "fmt"

func main() {
	fmt.Println(generate(5))

}
func generate(numRows int) [][]int {
	arr := make([][]int, numRows)
	arr[0] = []int{1, 0}
	for i := 1; i < numRows; i++ {
		arr[i] = make([]int, i+2)
		arr[i][0] = 1
		for j := 1; j <= i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}
	pascalArr := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		pascalArr[i] = arr[i][:i+1]
	}
	return pascalArr
}

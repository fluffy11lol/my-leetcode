package main

import (
	"fmt"
)

func main() {
	fmt.Println(getWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000))

}
func getWinner(arr []int, k int) int {
	n := len(arr)
	current := arr[0]
	winCount := 0

	for i := 1; i < n; i++ {
		if arr[i] > current {
			current = arr[i]
			winCount = 1
		} else {
			winCount++
		}

		if winCount == k {
			return current
		}
	}

	return current
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))

}
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	numTracker := 0
	minIncrement := 0

	for _, num := range nums {
		if numTracker < num {
			numTracker = num
		}
		minIncrement += numTracker - num
		numTracker++
	}

	return minIncrement
}

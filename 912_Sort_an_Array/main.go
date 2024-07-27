package main

import (
	"slices"
)

func main() {

}
func sortArray(nums []int) []int {
	// find min and max
	mi, mx := slices.Min(nums), slices.Max(nums)

	size := mx - mi + 1
	count := make([]int, size)
	for _, num := range nums {
		count[num-mi]++
	}

	index := 0
	for v, c := range count {
		for c > 0 {
			nums[index] = v + mi
			index++
			c--
		}
	}
	return nums
}

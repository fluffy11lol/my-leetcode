package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))

}
func countPairs(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l, r := 0, len(nums)-1
	count := 0

	for l < len(nums) {
		if nums[r]+nums[l] < target {
			count += r - l
			l++
		} else {
			r--
		}
		if l == r {
			break
		}
	}
	return count
}

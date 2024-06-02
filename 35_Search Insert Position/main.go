package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 6))
}
func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}
	l := 0
	r := len(nums) - 1
	m := (l + r) / 2
	for l <= r {
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
		m = (l + r) / 2
	}
	if nums[m] > target {
		return m
	} else {
		return m + 1
	}
}

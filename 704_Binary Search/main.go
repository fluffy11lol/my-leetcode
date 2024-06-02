package main

import "fmt"

func main() {
	fmt.Println(search([]int{1}, 1))
}

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
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
	return -1
}

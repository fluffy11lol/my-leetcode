package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))

}
func searchRange(nums []int, target int) []int {
	pos := []int{-1, -1}
	l := 0
	r := len(nums) - 1
	m := (l + r) / 2
	for l <= r {
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			i := m
			for ; i >= 0 && nums[i] == target; i-- {
			}
			i++
			pos[0] = i
			i = m
			for ; i < len(nums) && nums[i] == target; i++ {
			}
			i--
			pos[1] = i
			break
		}
		m = (l + r) / 2
	}
	return pos
}

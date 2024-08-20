package main

import "slices"

func main() {

}
func minMoves(nums []int) int {
	minNum := slices.Min(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i] - minNum
	}
	return res
}

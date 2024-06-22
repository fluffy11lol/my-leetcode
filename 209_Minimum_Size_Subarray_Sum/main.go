package main

import "math"

func main() {

}
func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	curSum := 0
	minL := math.MaxInt
	for ; r < len(nums); r++ {
		curSum += nums[r]
		for curSum >= target {
			if r-l+1 < minL {
				minL = r - l + 1
			}
			curSum -= nums[l]
			l++
		}
	}
	if minL < math.MaxInt {
		return minL
	} else {
		return 0
	}
}

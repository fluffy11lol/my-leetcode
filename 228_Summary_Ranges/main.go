package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(summaryRanges([]int{-1, 2}))

}
func summaryRanges(nums []int) []string {
	arrStr := make([]string, 0)
	if len(nums) == 0 {
		return arrStr
	}
	a, b := -4, -4
	for i := 0; i < len(nums); i++ {
		if a == -4 {
			a = nums[i]
			b = nums[i]
			continue
		}
		if nums[i] == b+1 {
			b = nums[i]
			continue
		}
		i--
		if a == b {
			arrStr = append(arrStr, strconv.Itoa(a))
			a = -4
		} else {
			arrStr = append(arrStr, strconv.Itoa(a)+"->"+strconv.Itoa(b))
			a = -4
		}
	}
	if a == b {
		arrStr = append(arrStr, strconv.Itoa(a))
	} else {
		arrStr = append(arrStr, strconv.Itoa(a)+"->"+strconv.Itoa(b))
	}
	return arrStr
}

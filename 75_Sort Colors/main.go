package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 2}
	sortColors(nums)
	fmt.Println(nums)
}
func sortColors(nums []int) {
	freqArr := make([]int, 3)
	for _, n := range nums {
		freqArr[n]++
	}
	i := 0
	for c := range freqArr {
		for range freqArr[c] {
			nums[i] = c
			i++
		}
	}
}

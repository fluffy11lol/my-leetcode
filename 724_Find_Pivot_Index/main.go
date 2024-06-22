package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{2}))

}
func pivotIndex(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	var leftSum int
	for i, n := range nums {
		if leftSum == sum-leftSum-n {
			return i
		}
		leftSum += n
	}
	return -1
}

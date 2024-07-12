package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))

}
func shuffle(nums []int, n int) []int {
	shNums := make([]int, 2*n)
	for i := 0; i < n; i++ {
		shNums[2*i] = nums[i]
		shNums[2*i+1] = nums[i+n]
	}
	return shNums
}

package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1}))

}
func missingNumber(nums []int) int {
	n := len(nums) + 1
	expSum := (n - 1) * n / 2
	actSum := 0
	for _, num := range nums {
		actSum += num
	}
	return expSum - actSum
}

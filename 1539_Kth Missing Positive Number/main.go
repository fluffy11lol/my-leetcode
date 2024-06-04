package main

import "fmt"

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))

}
func findKthPositive(nums []int, k int) int {
	i, num := 0, 1

	for k > 0 && i < len(nums) {
		if nums[i] > num {
			k--
		} else {
			i++
		}
		num++
	}
	return num - 1 + k
}

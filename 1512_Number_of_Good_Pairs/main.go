package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1, 1}))

}
func numIdenticalPairs(nums []int) int {
	numsMap := make(map[int]int)
	c := 0
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
	}
	for _, v := range numsMap {

		c += (v * (v - 1)) / 2

	}
	return c
}

package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))

}
func singleNumber(nums []int) int {
	var xor int
	for _, n := range nums {
		xor ^= n
	}
	return xor
}

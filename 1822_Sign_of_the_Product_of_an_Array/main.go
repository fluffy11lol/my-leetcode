package main

import "fmt"

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))

}
func arraySign(nums []int) int {
	negNums := 0
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			negNums++
		}
	}
	if negNums%2 == 1 {
		return -1
	}
	return 1
}

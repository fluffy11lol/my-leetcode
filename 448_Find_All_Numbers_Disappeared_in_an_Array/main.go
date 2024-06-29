package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers([]int{1, 1, 3}))

}
func findDisappearedNumbers(nums []int) []int {
	l := len(nums)
	presArr := make([]bool, l)
	for _, num := range nums {
		if num <= l {
			presArr[num-1] = true
		}
	}
	disapNums := make([]int, 0)
	for i := range presArr {
		if !presArr[i] {
			disapNums = append(disapNums, i+1)
		}
	}
	return disapNums
}

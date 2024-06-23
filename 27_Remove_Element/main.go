package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 1, 2, 1}
	fmt.Println(removeElement(arr, 1))
	fmt.Println(arr)

}
func removeElement(nums []int, val int) int {
	counter := 0
	for _, v := range nums {
		if v != val {
			nums[counter] = v
			counter++
		}
	}
	return counter
}

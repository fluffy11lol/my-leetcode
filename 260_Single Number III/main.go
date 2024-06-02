package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{-1, 0}))

}
func singleNumber(nums []int) []int {
	resultArr := []int{-199, -199}
	m := make(map[int]int, len(nums)/2+1)
	for _, n := range nums {
		m[n]++
	}
	for n, c := range m {
		if c == 1 {
			if resultArr[0] == -199 {
				resultArr[0] = n
			} else {
				resultArr[1] = n
				return resultArr
			}
		}
	}
	return resultArr
}

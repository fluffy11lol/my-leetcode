package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))

}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var j int
	var ok bool
	for i, n := range nums {
		if j, ok = m[target-n]; ok {
			return []int{i, j}
		}
		m[n] = i

	}
	return nil
}

package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 3, 31}, 2))

}
func containsNearbyDuplicate(nums []int, k int) bool {
	elemMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if j, ok := elemMap[nums[i]]; ok && i-j <= k {
			return true
		}
		elemMap[nums[i]] = i
	}
	return false
}

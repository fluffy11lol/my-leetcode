package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1, 2, 1}, 2))

}
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	var count, pSum int
	prefixSumMap := make(map[int]int, len(nums))
	// in case of pSum==k
	prefixSumMap[0] = 1
	for i := 0; i < len(nums); i++ {
		pSum += nums[i]
		count += prefixSumMap[pSum-k]
		prefixSumMap[pSum]++
	}
	return count
}

package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))

}
func numberOfSubarrays(nums []int, k int) int {
	prefixSumMap := make(map[int]int, len(nums))
	prefixSumMap[0] = 1
	pSum := 0
	count := 0
	for i := range nums {
		nums[i] &= 1 //odd ->1, even ->0
		pSum += nums[i]
		count += prefixSumMap[pSum-k]
		prefixSumMap[pSum]++
	}
	return count
}

//  s=2

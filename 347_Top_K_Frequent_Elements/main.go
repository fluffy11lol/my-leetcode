package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 2))

}
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for i := range nums {
		freqMap[nums[i]]++
	}
	freqMap2 := make(map[int][]int)
	for n, v := range freqMap {
		freqMap2[v] = append(freqMap2[v], n)
	}
	res := make([]int, k)
	for i := len(nums); k > 0; i-- {
		if n, ok := freqMap2[i]; ok {
			for j := range n {
				res[k-1] = n[j]
				k--
			}
		}
	}
	return res
}

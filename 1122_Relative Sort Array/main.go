package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(relativeSortArray([]int{1, 2, 3, 4, 0}, []int{1, 2, 3}))

}
func relativeSortArray(arr1 []int, arr2 []int) []int {
	freqMap := make(map[int]int, len(arr2))
	for _, n := range arr1 {
		freqMap[n]++
	}
	newArr := make([]int, len(arr1))
	i := 0
	for _, n := range arr2 {
		for range freqMap[n] {
			newArr[i] = n
			i++
		}
		freqMap[n] = 0
	}
	unsortedSliceI := i
	for n, f := range freqMap {
		if f != 0 {
			for range f {
				newArr[i] = n
				i++
			}
		}
	}
	sort.Slice(newArr[unsortedSliceI:], func(i, j int) bool {
		return newArr[i+unsortedSliceI] < newArr[j+unsortedSliceI]
	})
	return newArr
}

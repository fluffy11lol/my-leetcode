package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumImportance(5, [][]int{{0, 3}, {2, 4}, {1, 3}}))

}
func maximumImportance(n int, roads [][]int) int64 {
	impArr := make([]int, n)
	for _, road := range roads {
		impArr[road[0]]++
		impArr[road[1]]++
	}
	sort.Slice(impArr, func(i, j int) bool {
		return impArr[i] < impArr[j]
	})
	var maxImportance int64
	for i := range n {
		maxImportance += int64((i + 1) * impArr[i])
	}
	return maxImportance
}

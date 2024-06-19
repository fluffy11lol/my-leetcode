package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProfitAssignment([]int{68, 35, 52, 47, 86},
		[]int{67, 17, 1, 81, 3}, []int{92, 10, 85, 84, 82}))

}
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	jobs := make([][2]int, len(difficulty))
	for i := range jobs {
		jobs[i][0], jobs[i][1] = difficulty[i], profit[i]
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] > jobs[j][1]
	})
	maxProfit := 0

	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})
	curJobIndex := 0
	for _, wDif := range worker {
		for curJobIndex < len(jobs) {
			if jobs[curJobIndex][0] <= wDif {
				maxProfit += jobs[curJobIndex][1]
				break
			} else {
				curJobIndex++
			}
		}
		if curJobIndex >= len(jobs) {
			return maxProfit
		}
	}
	return maxProfit
}

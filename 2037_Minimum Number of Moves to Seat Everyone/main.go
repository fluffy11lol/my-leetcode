package main

import (
	"math"
	"sort"
)

func main() {

}
func minMovesToSeat(seats []int, students []int) int {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})
	sort.Slice(students, func(i, j int) bool {
		return students[i] < students[j]
	})
	movesCount := 0
	for i := range students {
		movesCount += int(math.Abs(float64(students[i] - seats[i])))
	}
	return movesCount
}

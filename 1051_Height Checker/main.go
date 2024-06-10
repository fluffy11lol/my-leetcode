package main

import "sort"

func main() {

}
func heightChecker(heights []int) int {
	sortHeights := make([]int, len(heights))

	copy(sortHeights, heights)

	sort.Ints(sortHeights)

	counter := 0

	for index, height := range sortHeights {
		if height != heights[index] {
			counter++
		}
	}

	return counter
}

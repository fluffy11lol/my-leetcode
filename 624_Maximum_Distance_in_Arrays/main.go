package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))

}
func maxDistance(arrays [][]int) int {
	minVal, minValIdx, maxVal := math.MaxInt32, -1, math.MinInt32
	minVal1, maxValIdx1, maxVal1 := math.MaxInt32, -1, math.MinInt32

	for idx, array := range arrays {
		minValueArr, maxValueArr := array[0], array[len(array)-1]
		if minValueArr < minVal {
			minVal = minValueArr
			minValIdx = idx
		}
		if maxValueArr > maxVal1 {
			maxVal1 = maxValueArr
			maxValIdx1 = idx
		}
	}

	for idx, array := range arrays {
		minValueArr, maxValueArr := array[0], array[len(array)-1]
		if maxValueArr > maxVal && idx != minValIdx {
			maxVal = maxValueArr
		}
		if minValueArr < minVal1 && idx != maxValIdx1 {
			minVal1 = minValueArr
		}
	}
	if abs(maxVal-minVal) < abs(maxVal1-minVal1) {
		return abs(maxVal1 - minVal1)
	}
	return abs(maxVal - minVal)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDays([]int{7, 7, 7, 7, 12, 7, 7}, 1, 3))

}
func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if n < m*k {
		return -1
	}
	low := 1
	maxi := 0
	for _, day := range bloomDay {
		maxi = max(maxi, day)
	}
	high := maxi
	for low <= high {
		mid := low + (high-low)/2
		count := 0
		bouq := 0
		flag := 0
		for _, day := range bloomDay {
			if day <= mid {
				count++
				if count == k {
					bouq++
					count = 0
				}
			} else {
				count = 0
			}
			if bouq == m {
				flag = 1
			}
		}
		if flag == 1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

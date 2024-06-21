package main

import "fmt"

func main() {
	fmt.Println(maxSatisfied([]int{8, 8}, []int{1, 0}, 2))

}
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	totalSatisfied := 0
	currentAdditional := 0
	maxAdditional := 0

	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			totalSatisfied += customers[i]
		}
	}

	for i := 0; i < n; i++ {
		currentAdditional += customers[i] * grumpy[i]
		if i >= minutes {
			currentAdditional -= customers[i-minutes] * grumpy[i-minutes]
		}
		if maxAdditional < currentAdditional {
			maxAdditional = currentAdditional
		}
	}
	return maxAdditional + totalSatisfied
}

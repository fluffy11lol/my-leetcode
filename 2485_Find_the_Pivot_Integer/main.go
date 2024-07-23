package main

import "fmt"

func main() {
	fmt.Println(pivotInteger(1))

}
func pivotInteger(n int) int {
	m := n / 2
	lSum, rSum := intervalSum(1, m), intervalSum(m, n)
	for lSum < rSum {
		rSum -= m
		m++
		lSum += m

	}
	if lSum == rSum {
		return m
	} else {
		return -1
	}
}
func intervalSum(a, b int) int {
	return (a + b) * (b - a + 1) / 2
}

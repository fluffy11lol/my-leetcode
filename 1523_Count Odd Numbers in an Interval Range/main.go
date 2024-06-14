package main

import "fmt"

// 1,2,3,4,5  1,2,3,4   2,3,4,5     2,3,4
func main() {
	fmt.Println(countOdds(0, 10))

}
func countOdds(low int, high int) int {
	if low%2 == 0 {
		low++
	}
	if low > high {
		return 0
	}
	return (high-low)/2 + 1
}

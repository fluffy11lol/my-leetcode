package main

import "fmt"

func main() {
	fmt.Println(countBits(2))

}
func countBits(n int) []int {
	n++
	binArr := make([]int, n)
	for i := 0; i < n; i++ {
		j := i
		for j > 0 {
			binArr[i] += j & 1
			j >>= 1
		}
	}
	return binArr
}

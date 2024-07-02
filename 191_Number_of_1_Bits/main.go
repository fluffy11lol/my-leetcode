package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(3))

}
func hammingWeight(n int) int {
	c := 0
	mask := 1
	for mask <= n {
		if n&mask > 0 {
			c++
		}
		mask <<= 1
	}
	return c
}

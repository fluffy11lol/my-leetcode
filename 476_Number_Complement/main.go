package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findComplement(5))

}
func findComplement(num int) int {
	pow := 0
	n := 1
	for num > n {
		n *= 2
		pow++
	}
	div := math.MaxInt >> (63 - pow)
	return (^num) & div
}

// alt: return (1<<(bits.Len(uint(num))) - 1) ^ num

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo(5))
}
func isPowerOfTwo(n int) bool {
	curPowOfTwo := 1
	for ; curPowOfTwo < n; curPowOfTwo <<= 1 {
	}
	return curPowOfTwo == n
}

package main

import "fmt"

func main() {
	fmt.Println(getSum(-12, -8))
}
func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

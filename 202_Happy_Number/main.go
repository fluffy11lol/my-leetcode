package main

import "fmt"

func main() {
	fmt.Println(isHappy(7))

}
func isHappy(n int) bool {
	sum, rem := 0, 0
	for n > 9 {
		sum = 0
		for n > 0 {
			rem = n % 10
			sum += rem * rem
			n /= 10
		}
		n = sum
	}
	if n == 1 || n == 7 {
		return true
	}
	return false
}

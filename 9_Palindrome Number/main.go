package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(122321))

}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversedX := 0
	for n := x; n > 0; {
		reversedX *= 10
		reversedX += n % 10
		n /= 10

	}
	return reversedX == x
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(121))

}
func mySqrt(num int) int {
	l := 0
	r := num + 1
	mid := (l + r) / 2
	sqrt := 0
	for l <= r {
		sqrt = mid * mid
		if sqrt > num {
			r = mid - 1
		} else if sqrt < num {
			l = mid + 1
		} else {
			return mid
		}
		mid = (l + r) / 2
	}
	return mid
}

package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(3))

}
func arrangeCoins(n int) int {
	l := 0
	r := n
	mid := (l + r) / 2
	for l <= r {
		sum := mid * (mid + 1) / 2
		if sum > n {
			r = mid - 1
		} else if sum < n {
			l = mid + 1
		} else {
			return mid
		}
		mid = (l + r) / 2
	}
	return 0
}

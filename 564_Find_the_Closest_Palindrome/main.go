package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(nearestPalindromic("100"))

}
func nearestPalindromic(n string) string {
	m := len(n)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1}
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])
	for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} {
		y := x
		if m&1 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}
	res := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		if candidate != selfNumber {
			if res == -1 ||
				abs(candidate-selfNumber) < abs(res-selfNumber) ||
				abs(candidate-selfNumber) == abs(res-selfNumber) &&
					candidate < res {
				res = candidate
			}
		}
	}
	return strconv.Itoa(res)
}

//123345

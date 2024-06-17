package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgeSquareSum(6))

}
func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	a := int(math.Sqrt(float64(c / 2)))
	b := int(math.Sqrt(float64(c)))
	for i := a; i <= b; i++ {
		sqrI := i * i
		j := int(math.Sqrt(float64(c - sqrI)))
		if sqrI+j*j == c {
			return true
		}
	}
	return false
}

//if n <= 0:
//    print("Impossible")
//else:
//    a = int((n//2)**.5)
//    b = int(n**.5) + 1
//    for i in range(a, b):
//        j = int((n - i*i)**.5)
//        if i*i + j*j == n:
//            print(i, j)
//            break
//    else:
//        print("Impossible")

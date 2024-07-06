package main

import "fmt"

func main() {
	fmt.Println(passThePillow(3, 2))

}
func passThePillow(n int, time int) int {
	var i int
	steps := time % (2 * (n - 1))
	if steps > (n - 1) {
		i = n - (steps - n + 1)
	} else {
		i = 1 + steps
	}
	return i
}

// 1 2 3 4

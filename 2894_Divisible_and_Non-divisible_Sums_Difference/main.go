package main

import "fmt"

func main() {

	fmt.Println(differenceOfSums(5, 6))
}
func differenceOfSums(n int, m int) int {
	d := (1 + n) * n / 2
	if n < m {
		return d
	}
	for i := m; i <= n; i += m {
		d -= 2 * i
	}
	return d
}

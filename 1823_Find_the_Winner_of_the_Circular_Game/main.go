package main

import "fmt"

func main() {
	fmt.Println(findTheWinner(3, 1))

}
func findTheWinner(n int, k int) int {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i + 1
	}
	losePlayer := 0
	for len(p) > 1 {
		losePlayer = (losePlayer + k - 1) % len(p)
		p = append(p[:losePlayer], p[losePlayer+1:]...)
		losePlayer = (losePlayer) % len(p)
	}
	return p[0]
}

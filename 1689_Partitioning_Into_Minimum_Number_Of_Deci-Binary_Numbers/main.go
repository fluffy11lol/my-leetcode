package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(minPartitions("82734"))

}
func minPartitions(n string) int {
	maxN := 0
	for i := len(n) - 1; i >= 0; i-- {
		cur, _ := strconv.Atoi(string(n[i]))
		if cur > maxN {
			maxN = cur
		}
		if maxN == 9 {
			break
		}
	}
	return maxN
}

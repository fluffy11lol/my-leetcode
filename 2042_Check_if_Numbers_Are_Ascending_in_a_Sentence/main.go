package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(areNumbersAscending("33 box has 3 blue 4 red 6 green and 12 yellow marbles"))

}
func areNumbersAscending(s string) bool {
	tokensArr := make([]int, 0)
	tokensArr = append(tokensArr, -1)
	sArr := strings.Split(s, " ")
	curI := 0
	for i := range sArr {
		n, err := strconv.Atoi(sArr[i])
		if err != nil {
			continue
		}
		tokensArr = append(tokensArr, n)
		curI++
		if tokensArr[curI] <= tokensArr[curI-1] {
			return false
		}
	}
	return true
}

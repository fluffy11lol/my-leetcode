package main

import (
	"strings"
)

func main() {

}
func findWordsContaining(words []string, x byte) []int {
	cArr := make([]int, 0)
	for i := range words {
		if strings.Contains(words[i], string(x)) {
			cArr = append(cArr, i)
		}
	}
	return cArr
}

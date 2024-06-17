package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mergeAlternately("1232", "234"))

}
func mergeAlternately(word1 string, word2 string) string {
	w1Arr := []byte(word1)
	w2Arr := []byte(word2)
	mergedArr := make([]byte, len(w2Arr)+len(w1Arr))
	p1, p2 := 0, 0
	for range int(math.Min(float64(len(w1Arr)), float64(len(w2Arr)))) {
		mergedArr[p1+p2] = w1Arr[p1]
		p1++
		mergedArr[p1+p2] = w2Arr[p2]
		p2++
	}
	if len(w1Arr) > p1 {
		for ; p1 < len(w1Arr); p1++ {
			mergedArr[p1+p2] = w1Arr[p1]
		}
	}
	if len(w2Arr) > p2 {
		for ; p2 < len(w2Arr); p2++ {
			mergedArr[p1+p2] = w2Arr[p2]
		}
	}
	return string(mergedArr)
}

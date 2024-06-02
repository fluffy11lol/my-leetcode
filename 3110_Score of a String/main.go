package main

import "fmt"

func main() {
	fmt.Println(scoreOfString("zaz"))

}
func scoreOfString(s string) int {
	asciiArr := []byte(s)
	var score int
	var cur = asciiArr[0]
	var next uint8
	for i := 1; i < len(s); i++ {
		next = asciiArr[i]
		if cur > next {
			score += int(cur - next)
		} else {
			score += int(next - cur)
		}
		cur = next
	}
	return score
}

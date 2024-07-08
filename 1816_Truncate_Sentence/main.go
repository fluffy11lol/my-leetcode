package main

import "fmt"

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))

}
func truncateSentence(s string, k int) string {
	spacesCount := 0
	i := 0
	for ; spacesCount < k && i < len(s); i++ {
		if s[i] == ' ' {
			spacesCount++
		}
	}
	if i == len(s) {
		return s[:i]
	}
	return s[:i-1]
}

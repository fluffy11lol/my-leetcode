package main

import "fmt"

func main() {
	fmt.Println(appendCharacters("abcde", "a"))

}
func appendCharacters(s string, t string) int {
	l1 := 0
	maxN := 0
	for l1 < len(s) && maxN < len(t) {
		if s[l1] == t[maxN] {
			maxN++
		}
		l1++
	}
	return len(t) - maxN
}

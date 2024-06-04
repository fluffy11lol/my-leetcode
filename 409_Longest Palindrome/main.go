package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("a"))

}
func longestPalindrome(s string) int {
	arr := [128]int{}
	for _, ch := range s {
		arr[ch]++
	}
	length := 0
	oddFound := false
	for _, c := range arr {
		if c > 1 {
			length += (c / 2) * 2
		}
		if !oddFound && c%2 == 1 {
			length++
			oddFound = true
		}
	}
	return length
}

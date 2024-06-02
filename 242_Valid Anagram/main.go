package main

import "fmt"

func main() {
	fmt.Println(isAnagram("123", "123"))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := make([]int32, 255)
	for i := range s {
		a[s[i]]++
		a[t[i]]--
	}
	for _, n := range a {
		if n != 0 {
			return false
		}
	}

	return true
}

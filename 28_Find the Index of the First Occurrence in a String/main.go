package main

import "fmt"

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))

}
func strStr(haystack string, needle string) int {
	for i, v := range haystack {
		if len(haystack)-i < len(needle) {
			return -1
		}
		if string(v) == string(needle[0]) {
			r := 0
			for ; r < len(needle); r++ {
				if haystack[i+r] != needle[r] {
					break
				}
			}
			if r == len(needle) {
				return i
			}
		}
	}
	return -1
}

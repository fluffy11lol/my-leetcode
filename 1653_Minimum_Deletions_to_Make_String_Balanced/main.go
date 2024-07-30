package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumDeletions("bbaaaaabb"))

}
func minimumDeletions(s string) int {
	var res, bCount int

	for i := range s {
		if s[i] == 'a' {
			res = min(res+1, bCount)
		} else {
			bCount++
		}
	}
	return res
}

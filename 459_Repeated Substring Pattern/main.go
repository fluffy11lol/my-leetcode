package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(repeatedSubstringPattern("aa"))

}
func repeatedSubstringPattern(s string) bool {
	sArr := []byte(s)
	l := len(s)
DiapCheck:
	for i := 1; i <= l/2; i++ {
		if l%i == 0 {
			t := l / i
			j := i
			for range t - 1 {
				if bytes.Compare(sArr[0:i], sArr[j:j+i]) != 0 {
					continue DiapCheck
				}
				j += i
			}
			return true
		}
	}
	return false
}

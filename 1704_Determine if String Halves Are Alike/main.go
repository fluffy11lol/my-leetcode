package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("booikk"))

}
func halvesAreAlike(s string) bool {
	sum := 0
	halfLen := len(s) / 2
	for i := 0; i < halfLen; i++ {
		ch1 := s[i]
		ch2 := s[halfLen+i]
		switch ch1 {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			sum++

		}
		switch ch2 {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			sum--

		}
	}
	if sum != 0 {
		return false
	}
	return true
}

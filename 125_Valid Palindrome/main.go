package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	l := 0
	r := len(s) - 1

	for l < r {
		for !isAlphaNumeric(runes[l]) && l < r {
			l++
		}

		for !isAlphaNumeric(runes[r]) && l < r {
			r--
		}

		if l < r && unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]) {
			return false
		} else {
			l++
			r--
		}
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
